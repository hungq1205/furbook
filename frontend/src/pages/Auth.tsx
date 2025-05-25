import React, { useState } from 'react';
import { motion } from 'framer-motion';
import { useNavigate } from 'react-router-dom';
import { PawPrint, User2, UserRoundSearch, Lock, LockKeyhole, AlertCircle } from 'lucide-react';
import Button from '../components/common/Button';
import Card from '../components/common/Card';
import { useAuth } from '../services/authService';
import { HttpError } from '../services/util';

const Auth: React.FC = () => {
  const authService = useAuth();
  const navigate = useNavigate();

  const [isLogin, setIsLogin] = useState(true)
  const [username, setUsername] = useState('');
  const [displayName, setDisplayName] = useState('');
  const [password, setPassword] = useState('');
  const [cfPassword, setCfPassword] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    try {
        if (isLogin) {
            await authService.login({ username, password });
        } else {
            await authService.signup({ username, password, displayName });
        }
        navigate('/');
    } catch (err) {
        setError(err instanceof HttpError ? err.message : 'Failed to sign up');
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-md">
        <motion.div
          initial={{ scale: 0.5, opacity: 0 }}
          animate={{ scale: 1, opacity: 1 }}
          transition={{ duration: 0.5 }}
          className="flex justify-center"
        >
          <PawPrint className="text-primary-600" size={48} />
        </motion.div>
        <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
          Join PetPals
        </h2>
        {
            isLogin ?
            (
                <p className="mt-2 text-center text-sm text-gray-600">
                    Don't have an account?&nbsp;
                    <span onClick={() => setIsLogin(false)} className="font-medium text-primary-600 hover:text-primary-500 cursor-pointer">
                        Sign up
                    </span>
                </p>
            ) :
            (
                <p className="mt-2 text-center text-sm text-gray-600">
                    Already have an account?&nbsp;
                    <span onClick={() => setIsLogin(true)} className="font-medium text-primary-600 hover:text-primary-500 cursor-pointer">
                        Sign in
                    </span>
                </p>
            )
        }
      </div>

      <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
        <Card>
          <form onSubmit={handleSubmit} className="space-y-6 p-8">
            {error && (
              <div className="flex items-center space-x-2 text-error-600 bg-error-50 p-3 rounded-md">
                <AlertCircle size={20} />
                <span>{error}</span>
              </div>
            )}

            <div>
              <label htmlFor="username" className="block text-sm font-medium text-gray-700">
                Username
              </label>
              <div className="mt-1 relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <User2 className="h-5 w-5 text-gray-400" />
                </div>
                <input
                  id="username"
                  name="username"
                  type="text"
                  autoComplete="username"
                  required
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                  className="appearance-none block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-primary-500 focus:border-primary-500"
                  placeholder="Username"
                />
              </div>
            </div>

            { !isLogin && 
            <div>
              <label htmlFor="displayName" className="block text-sm font-medium text-gray-700">
                Display Name
              </label>
              <div className="mt-1 relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <UserRoundSearch className="h-5 w-5 text-gray-400" />
                </div>
                <input
                  id="displayName"
                  name="displayName"
                  type="displayName"
                  autoComplete="displayName"
                  required
                  value={displayName}
                  onChange={(e) => setDisplayName(e.target.value)}
                  className="appearance-none block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-primary-500 focus:border-primary-500"
                  placeholder="Display name"
                />
              </div>
            </div> 
            }

            <div>
              <label htmlFor="password" className="block text-sm font-medium text-gray-700">
                Password
              </label>
              <div className="mt-1 relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <Lock className="h-5 w-5 text-gray-400" />
                </div>
                <input
                  id="password"
                  name="password"
                  type="password"
                  required
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  className="appearance-none block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-primary-500 focus:border-primary-500"
                  placeholder="Password"
                />
              </div>
            </div>

            { !isLogin && 
            <div>
              <label htmlFor="cf-password" className="block text-sm font-medium text-gray-700">
                Confirm Password
              </label>
              <div className="mt-1 relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <LockKeyhole className="h-5 w-5 text-gray-400" />
                </div>
                <input
                  id="cf-password"
                  name="cf-password"
                  type="cf-password"
                  required
                  value={cfPassword}
                  onChange={(e) => setCfPassword(e.target.value)}
                  className="appearance-none block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-primary-500 focus:border-primary-500"
                  placeholder="Password"
                />
              </div>
            </div>
            }

            <div>
              <Button
                type="submit"
                variant="primary"
                fullWidth
              >
                { isLogin ? "Sign in" : "Sign up" }
              </Button>
            </div>
          </form>
        </Card>
      </div>
    </div>
  );
};

export default Auth;