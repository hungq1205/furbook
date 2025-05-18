import { User } from '../types/user';
import { defaultAuthHeaders, BASE_URL, HttpError } from './util';

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
      body: JSON.stringify({ friend: username })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
  },

  async checkFriendship(username: string): Promise<{ isFriend: boolean }> {
    const response = await fetch(`${USER_URL}/check-friendship`, {
      method: 'POST',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ username })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
    return response.json();
  },

  async checkFriendRequest(username: string): Promise<{ exists: boolean, type: 'sent' | 'received' | null }> {
    const response = await fetch(`${USER_URL}/check-friend-request`, {
      method: 'POST',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ username })
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