import { Friendship, User } from '../types/user';
import { BASE_URL, defaultAuthHeaders, HttpError } from './util';

const USER_URL = `${BASE_URL}/user`;

export const userService = {
  async getUser(username: string): Promise<User> {
    const response = await fetch(`${USER_URL}/${username}`, {
      headers: defaultAuthHeaders()
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
    return response.json();
  },

  async getUsers(usernames: string[]): Promise<User[]> {
    const response = await fetch(`${USER_URL}/list`, {
      method: 'POST',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ usernames })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
    return response.json();
  },

  async getFriends(): Promise<User[]> {
    const response = await fetch(`${USER_URL}/friends`, {
      headers: defaultAuthHeaders()
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
    return response.json();
  },

  async getFriendRequests(): Promise<User[]> {
    const response = await fetch(`${USER_URL}/friend-requests`, {
      headers: defaultAuthHeaders()
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
    return response.json();
  },

  async sendFriendRequest(username: string): Promise<void> {
    const response = await fetch(`${USER_URL}/friend-requests`, {
      method: 'POST',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ receiver: username })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
  },

  async checkFriendship(username: string): Promise<{ friendship: Friendship }> {
    const response = await fetch(`${USER_URL}/check-friendship/${username}`, {
      method: 'GET',
      headers: defaultAuthHeaders(),
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
    return response.json();
  },

  async removeFriend(username: string): Promise<void> {
    const response = await fetch(`${USER_URL}/friends`, {
      method: 'DELETE',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ friend: username })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
  },

  async declineFriendRequest(username: string): Promise<void> {
    const response = await fetch(`${USER_URL}/friend-requests/decline`, {
      method: 'DELETE',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ sender: username })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
  },

  async revokeFriendRequest(username: string): Promise<void> {
    const response = await fetch(`${USER_URL}/friend-requests/revoke`, {
      method: 'DELETE',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ receiver: username })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
  }
};