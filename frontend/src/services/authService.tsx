import {
  createContext,
  useContext,
  useEffect,
  useState,
  ReactNode,
  useCallback,
} from 'react';
import { User } from '../types/user';
import { userService } from './userService';
import { BASE_URL, defaultAuthHeaders, defaultHeaders, HttpError } from './util';

const AUTH_URL = `${BASE_URL}/auth`;

interface LoginPayload {
  username: string;
  password: string;
}

interface SignupPayload {
  username: string;
  displayName: string;
  password: string;
}

interface AuthContextValue {
  token: string | null;
  currentUser: User | null;
  currentUserFriends: User[];
  isAuthenticated: boolean;
  login: (data: LoginPayload) => Promise<void>;
  signup: (data: SignupPayload) => Promise<void>;
  logout: () => void;
  refresh: () => Promise<void>;
}

export const AuthContext = createContext<AuthContextValue | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [token, setToken] = useState<string | null>(() => localStorage.getItem('token'));
  const [currentUser, setCurrentUser] = useState<User | null>(null);
  const [currentUserFriends, setCurrentUserFriends] = useState<User[]>([]);

  const isAuthenticated = !!token;

  const saveToken = (newToken: string) => {
    localStorage.setItem('token', newToken);
    setToken(newToken);
  };

  const clearToken = () => {
    localStorage.removeItem('token');
    setToken(null);
  };

  const refresh = useCallback(async () => {
    if (!token) return;
    const response = await fetch(`${AUTH_URL}/check`, { headers: defaultAuthHeaders()});
    if (!response.ok) throw new HttpError(response.status, await response.json());

    const { token: newToken, user } = await response.json() as { token: string; user: User };
    saveToken(newToken);
    setCurrentUser(user);

    const friends = await userService.getFriends();
    setCurrentUserFriends(friends);
  }, [token]);

  const login = async (data: LoginPayload) => {
    const response = await fetch(`${AUTH_URL}/login`, {
      method: 'POST',
      headers: defaultHeaders,
      body: JSON.stringify(data),
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());

    const { token: newToken, user } = await response.json() as { token: string; user: User };
    saveToken(newToken);
    setCurrentUser(user);

    const friends = await userService.getFriends();
    setCurrentUserFriends(friends);
  };

  const signup = async (data: SignupPayload) => {
    const response = await fetch(`${AUTH_URL}/signup`, {
      method: 'POST',
      headers: defaultHeaders,
      body: JSON.stringify(data),
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
  };

  const logout = () => {
    clearToken();
    setCurrentUser(null);
    setCurrentUserFriends([]);
  };
  
  return (
    <AuthContext.Provider value={{ 
      token,
      currentUser, 
      currentUserFriends, 
      isAuthenticated, 
      login, 
      signup, 
      logout, 
      refresh 
    }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = (): AuthContextValue => {
  const context = useContext(AuthContext);
  if (!context) throw new Error('useAuth must be used within an AuthProvider');
  return context;
};
