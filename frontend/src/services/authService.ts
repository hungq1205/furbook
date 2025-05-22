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

class AuthService {
  private static instance: AuthService;
  private token: string | null = null;
  private currentUser: User = {} as User;
  private currentUserFriends: User[] = [];

  private constructor() {
    this.token = localStorage.getItem('token');
  }

  static getInstance(): AuthService {
    if (!AuthService.instance) {
      AuthService.instance = new AuthService();
    }
    return AuthService.instance;
  }

  async check() {
    const response = await fetch(`${AUTH_URL}/check`, { headers: defaultAuthHeaders() });
    if (!response.ok) throw new HttpError(response.status, await response.json());

    const { token, user } = await response.json() as { token: string; user: User };
    this.setToken(token);
    this.currentUser = user;

    const friends = await userService.getFriends();
    this.currentUserFriends = friends;
  }

  async signup(data: SignupPayload): Promise<void> {
    const response = await fetch(`${AUTH_URL}/signup`, {
      method: 'POST',
      headers: defaultHeaders,
      body: JSON.stringify(data)
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
  }

  async login(data: LoginPayload): Promise<User> {
    const response = await fetch(`${AUTH_URL}/login`, {
      method: 'POST',
      headers: defaultHeaders,
      body: JSON.stringify(data)
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());

    const { token, user } = await response.json() as { token: string; user: User };
    this.setToken(token);
    this.currentUser = user;

    const friends = await userService.getFriends();
    this.currentUserFriends = friends;
    return user;
  }

  logout(): void {
    this.token = null;
    this.currentUser = {} as User;
    this.currentUserFriends = [];
    localStorage.removeItem('token');
  }

  getToken(): string | null {
    return this.token;
  }

  private setToken(token: string): void {
    this.token = token;
    localStorage.setItem('token', token);
  }

  getCurrentUser(): User {
    return this.currentUser;
  }

  getCurrentUserFriends(): User[] {
    return this.currentUserFriends;
  }

  isAuthenticated(): boolean {
    console.log('isAuthenticated', this.token);
    return !!this.token;
  }

  getAuthHeaders(): HeadersInit {
    return this.token
      ? { Authorization: `Bearer ${this.token}` }
      : {};
  }
}

export const authService = AuthService.getInstance();