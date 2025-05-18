import { Message } from '../types/message';
import { defaultAuthHeaders, BASE_URL, HttpError } from './util';

const MESSAGE_URL = `${BASE_URL}/message`;

export const messageService = {
  async getDirectMessages(username: string): Promise<Message[]> {
    const response = await fetch(`${MESSAGE_URL}/direct?oppUsername=${encodeURIComponent(username)}`, {
      headers: defaultAuthHeaders()
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
    return response.json();
  },

  async getGroupMessages(groupId: number): Promise<Message[]> {
    const response = await fetch(`${MESSAGE_URL}/group/${groupId}`, {
      headers: defaultAuthHeaders()
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
    return response.json();
  },

  async sendDirectMessage(username: string, content: string): Promise<Message> {
    const response = await fetch(`${MESSAGE_URL}/direct`, {
      method: 'POST',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ oppUsername: username, content: content })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
    return response.json();
  },

  async sendGroupMessage(groupId: number, content: string): Promise<Message> {
    const response = await fetch(`${MESSAGE_URL}/group/${groupId}`, {
      method: 'POST',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ content })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());;
    return response.json();
  }
};