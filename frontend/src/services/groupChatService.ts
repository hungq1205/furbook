import { User } from '../types/user';
import { GroupChat } from '../types/message';
import { BASE_URL, defaultAuthHeaders, HttpError } from './util';

const GROUP_CHAT_URL = `${BASE_URL}/group-chat`;

export const groupChatService = {
  async getGroups(): Promise<GroupChat[]> {
    const response = await fetch(`${GROUP_CHAT_URL}`, {
      headers: defaultAuthHeaders()
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
    return response.json();
  },

  async getGroupDetails(groupId: number): Promise<GroupChat> {
    const response = await fetch(`${GROUP_CHAT_URL}/${groupId}`, {
      headers: defaultAuthHeaders()
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
    return response.json();
  },

  async getGroupMembers(groupId: number): Promise<User[]> {
    const response = await fetch(`${GROUP_CHAT_URL}/${groupId}/members`, {
      headers: defaultAuthHeaders()
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
    return response.json();
  },

  async createGroup(ownername: string, groupName: string, members: string[]): Promise<GroupChat> {
    const response = await fetch(`${GROUP_CHAT_URL}`, {
      method: 'POST',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ ownername, groupName, members })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
    return response.json();
  },

  async addMember(groupId: number, username: string): Promise<void> {
    const response = await fetch(`${GROUP_CHAT_URL}/${groupId}/members`, {
      method: 'POST',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ username })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
  },

  async removeMember(groupId: number, username: string): Promise<void> {
    const response = await fetch(`${GROUP_CHAT_URL}/${groupId}/members`, {
      method: 'DELETE',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ username })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
  }
};