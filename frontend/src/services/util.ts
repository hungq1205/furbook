export const BASE_URL = (import.meta.env.VITE_API_URL || 'http://localhost:3000') + '/api';

export class HttpError extends Error {
  constructor(public status: number, message: string) {
    super(message);
    this.name = 'HttpError';
  }
}

export const defaultHeaders = {
    'Content-Type': 'application/json'
};

export const defaultAuthHeaders = () => ({
  'Content-Type': 'application/json',
  'Authorization': `Bearer ${localStorage.getItem('token')}`,
});