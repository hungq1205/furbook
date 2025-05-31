import { Notification } from '../types/notification';
import { defaultAuthHeaders, BASE_URL, HttpError } from './util';

const NOTI_URL = `${BASE_URL}/noti`;

export interface NotiPayload {
  icon: string,
  desc: string,
  link: string,
}

export const notiService = {
  async getById(notiId: string): Promise<Notification> {
    const response = await fetch(`${NOTI_URL}/${notiId}`, {headers: defaultAuthHeaders()}); 
    if (!response.ok) throw new HttpError(response.status, await response.json());
    return response.json();
  },

  async getByUsername(page: number): Promise<Notification[]> {
    const url = `${NOTI_URL}?page=${page}`;
    const response = await fetch(url, {headers: defaultAuthHeaders()});
    if (!response.ok) throw new HttpError(response.status, await response.json());
    return response.json();
  },

  async createNoti(payload: NotiPayload): Promise<Notification> {
    const response = await fetch(`${NOTI_URL}`, {
      method: 'POST',
      headers: defaultAuthHeaders(),
      body: JSON.stringify(payload)
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
    return response.json();
  },

  async markRead(notiId: string, read: boolean): Promise<Notification> {
    const response = await fetch(`${NOTI_URL}/${notiId}`, {
      method: 'PATCH',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ read })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
    return response.json();
  },

  async delete(notiId: string): Promise<void> {
    const response = await fetch(`${NOTI_URL}/${notiId}`, {
      method: 'DELETE',
      headers: defaultAuthHeaders()
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
  },
};