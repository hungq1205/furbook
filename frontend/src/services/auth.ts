import { User } from '../types/user';

const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:3000/api';

interface AuthResponse {
  token: string;
  user: User;
}

interface SignUpData {
  email: string;
  password: string;
  name: string;
}

interface SignInData {
  email: string;
  password: string;
}

class AuthService {
  private static instance: AuthService;
  private token: string | null = null;
  private currentUser: User | null = null;

  private constructor() {
    // Load token from localStorage
    this.token = localStorage.getItem('token');
  }

  static getInstance(): AuthService {
    if (!AuthService.instance) {
      AuthService.instance = new AuthService();
    }
    return AuthService.instance;
  }

  async signUp(data: SignUpData): Promise<User> {
    const response = await fetch(`${BASE_URL}/auth/signup`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.message || 'Failed to sign up');
    }

    const { token, user } = await response.json() as AuthResponse;
    this.setToken(token);
    this.currentUser = user;
    return user;
  }

  async signIn(data: SignInData): Promise<User> {
    const response = await fetch(`${BASE_URL}/auth/signin`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.message || 'Failed to sign in');
    }

    const { token, user } = await response.json() as AuthResponse;
    this.setToken(token);
    this.currentUser = user;
    return user;
  }

  signOut(): void {
    this.token = null;
    this.currentUser = null;
    localStorage.removeItem('token');
  }

  getToken(): string | null {
    return this.token;
  }

  private setToken(token: string): void {
    this.token = token;
    localStorage.setItem('token', token);
  }

  getCurrentUser(): User | null {
    return this.currentUser;
  }

  isAuthenticated(): boolean {
    return !!this.token;
  }

  // Add this to all API requests
  getAuthHeaders(): HeadersInit {
    return this.token
      ? { Authorization: `Bearer ${this.token}` }
      : {};
  }
}

export const authService = AuthService.getInstance();