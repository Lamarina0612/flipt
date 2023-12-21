import {
  faGithub,
  faGitlab,
  faGoogle,
  faOpenid,
  IconDefinition
} from '@fortawesome/free-brands-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { toLower, upperFirst } from 'lodash';
import { useMemo } from 'react';
import { Navigate } from 'react-router-dom';
import { useListAuthProvidersQuery } from '~/app/auth/authApi';
import logoFlag from '~/assets/logo-flag.png';
import Loading from '~/components/Loading';
import { NotificationProvider } from '~/components/NotificationProvider';
import ErrorNotification from '~/components/notifications/ErrorNotification';
import { useError } from '~/data/hooks/error';
import { useSession } from '~/data/hooks/session';
import { IAuthMethod } from '~/types/Auth';
import { useState } from 'react';

interface ILoginProvider {
  displayName: string;
  icon: IconDefinition;
}

interface IAuthDisplay {
  name: string;
  authorize_url: string;
  callback_url: string;
  icon: IconDefinition;
}

interface ILoginFormState {
  username: string;
  password: string;
}

const knownProviders: Record<string, ILoginProvider> = {
  google: {
    displayName: 'Google',
    icon: faGoogle
  },
  gitlab: {
    displayName: 'GitLab',
    icon: faGitlab
  },
  auth0: {
    displayName: 'Auth0',
    icon: faOpenid
  },
  github: {
    displayName: 'Github',
    icon: faGithub
  }
};

function InnerLoginButtons() {
  const { setError, clearError } = useError();

  const authorize = async (uri: string) => {
    const res = await fetch(uri, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    });

    if (!res.ok || res.status !== 200) {
      const { message } = await res.json();
      setError('Unable to authenticate: ' + message);
      return;
    }

    clearError();
    const body = await res.json();
    window.location.href = body.authorizeUrl;
  };
  const {
    data: listAuthProviders,
    isLoading,
    error
  } = useListAuthProvidersQuery();
  if (error) {
    setError(error);
  }

  const providers = useMemo(() => {
    return (listAuthProviders?.methods || [])
      .filter(
        (m: IAuthMethod) =>
          (m.method === 'METHOD_OIDC' || m.method === 'METHOD_GITHUB') &&
          m.enabled
      )
      .flatMap<any, IAuthDisplay>((m: IAuthMethod) => {
        if (m.method === 'METHOD_GITHUB') {
          return {
            name: 'Github',
            authorize_url: m.metadata.authorize_url,
            callback_url: m.metadata.callback_url,
            icon: faGithub
          };
        }
        if (m.method === 'METHOD_OIDC') {
          return Object.entries(m.metadata.providers).map(([k, value]) => {
            k = toLower(k);
            const v = value as {
              authorize_url: string;
              callback_url: string;
            };
            return {
              name: knownProviders[k]?.displayName || upperFirst(k), // if we dont know the provider, just capitalize the first letter
              authorize_url: v.authorize_url,
              callback_url: v.callback_url,
              icon: knownProviders[k]?.icon || faOpenid // if we dont know the provider icon, use the openid icon
            };
          });
        }
      });
  }, [listAuthProviders]);

  if (isLoading) {
    return <Loading />;
  }

  return (
    <>
      {providers.length > 0 && (
        <div className="mt-6 flex flex-col space-y-5">
          {providers.map((provider) => (
            <div key={provider.name}>
              <a
                href="#"
                className="bg-white text-gray-500 border-gray-300 inline-flex w-full justify-center rounded-md border px-4 py-2 text-sm font-medium shadow-sm hover:text-violet-500 hover:shadow-violet-300"
                onClick={(e) => {
                  e.preventDefault();
                  authorize(provider.authorize_url);
                }}
              >
                <span className="sr-only">Sign in with {provider.name}</span>
                <FontAwesomeIcon
                  icon={provider.icon}
                  className="text-gray h-5 w-5"
                  aria-hidden={true}
                />
                <span className="ml-2">With {provider.name}</span>
              </a>
            </div>
          ))}
        </div>
      )}
      {providers.length === 0 && (
        <div className="bg-white shadow sm:rounded-lg">
          <div className="px-4 py-5 sm:p-6">
            <h3 className="text-gray-900 text-base font-semibold leading-6">
              No Providers
            </h3>
            <div className="text-gray-500 mt-2 max-w-xl text-sm">
              <p>
                Authentication is set to{' '}
                <span className="font-medium">required</span>, however, there
                are no login providers configured. Please see the documentation
                for more information.
              </p>
            </div>
            <div className="mt-3 text-sm leading-6">
              <a
                href="https://www.flipt.io/docs/configuration/authentication#method-oidc"
                className="text-violet-600 font-semibold hover:text-violet-500"
              >
                Configuring Authentication
                <span aria-hidden="true"> &rarr;</span>
              </a>
            </div>
          </div>
        </div>
      )}
    </>
  );
}

function InnerLogin() {
  const { session } = useSession();
  const [showBasicAuth, setShowBasicAuth] = useState(false);
  const [loginForm, setLoginForm] = useState<ILoginFormState>({
    username: '',
    password: ''
  });

  const handleBasicAuthClick = () => {
    setShowBasicAuth(true);
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setLoginForm({ ... loginForm, [e.target.name]: e.target.value });
  };

  const handleSubmit = async(e: React.FormEvent) => {
    e.preventDefault();
  }

  if (showBasicAuth) {
    return (
      <div className="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
        <div className="sm:mx-auto sm:w-full sm:max-w-md">
          <img
            src={logoFlag}
            alt="logo"
            className="mx-auto h-20 w-auto"
          />
          <h2 className="mt-10 text-center text-3xl font-bold leading-9 tracking-tight text-gray-900">
            Login to Flipt
          </h2>
        </div>

        <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-md">
          <form className="space-y-6" onSubmit={handleSubmit}>
            <div>
              <label htmlFor="username" className="block text-sm font-medium leading-6 text-gray-900">
                Username
              </label>
              <div className="mt-2">
                <input
                  id="username"
                  name="username"
                  type="text"
                  autoComplete="username"
                  required
                  className="block w-full rounded-md border-purple-300 shadow-sm focus:border-purple-500 focus:ring focus:ring-purple-200"
                  value={loginForm.username}
                  onChange={handleInputChange}
                />
              </div>
            </div>

            <div>
              <label htmlFor="password" className="block text-sm font-medium leading-6 text-gray-900">
                Password
              </label>
              <div className="mt-2">
                <input
                  id="password"
                  name="password"
                  type="password"
                  autoComplete="current-password"
                  required
                  className="block w-full rounded-md border-purple-300 shadow-sm focus:border-purple-500 focus:ring focus:ring-purple-200"
                  value={loginForm.password}
                  onChange={handleInputChange}
                />
              </div>
            </div>

            <div className="flex items-center justify-between space-x-4">
              <button
                type="button"
                className="inline-flex justify-center rounded-md border border-transparent bg-gray-200 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-300 focus:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:ring-indigo-500"
                onClick={() => setShowBasicAuth(false)} 
              >
                Cancel
              </button>
              <button
                type="submit"
                className="inline-flex justify-center rounded-md border border-transparent bg-purple-600 px-4 py-2 text-sm font-medium text-white hover:bg-purple-700 focus:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:ring-indigo-500"
              >
                Submit
              </button>
            </div>
          </form>
        </div>
      </div>
    );
  }


  if (session && (!session.required || session.authenticated)) {
    return <Navigate to="/" />;
  }

  return (
    <>
      <div className="bg-white flex min-h-screen flex-col justify-center sm:px-6 lg:px-8">
        <main className="flex px-6 py-10">
          <div className="w-full overflow-x-auto px-4 sm:px-6 lg:px-8">
            <div className="sm:mx-auto sm:w-full sm:max-w-md">
              <img
                src={logoFlag}
                alt="logo"
                width={512}
                height={512}
                className="m-auto h-20 w-20"
              />
              <h2 className="text-gray-900 mt-6 text-center text-3xl font-bold tracking-tight">
                Login to Flipt
              </h2>
          
            </div>
            <div className="mt-8 max-w-sm sm:mx-auto sm:w-full md:max-w-lg">
              <div className="px-4 py-8 sm:px-10 flex justify-center">
              {!showBasicAuth && (
              <button
              className="bg-transparent hover:bg-purple-500 text-purple-700 font-semibold hover:text-white py-2 px-4 border border-purple-500 hover:border-transparent rounded"
              onClick={handleBasicAuthClick}
            >
              With Basic Authentication
            </button>
              )}
              </div>
            </div>
          </div>
        </main>
      </div>
    </>
  );
}

export default function Login() {
  return (
    <NotificationProvider>
      <InnerLogin />
      <ErrorNotification />
    </NotificationProvider>
  );
}
