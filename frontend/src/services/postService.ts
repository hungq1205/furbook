import { Post, Comment } from '../types/post';
import { defaultAuthHeaders, BASE_URL, HttpError } from './util';

const POST_URL = `${BASE_URL}/post`;

export type Media = {
    url: string;
    type: 'image' | 'video';
};

export type Location = {
    lat: number;
    lng: number;
};

export interface BlogPostPayload {
    content: string;
    medias: Media[];
}

export interface LostPostPayload extends BlogPostPayload {
    type: 'lost' | 'found';
    contactInfo: string;
    lostAt?: string;
    area: Location;
    lastSeen: Location;
}

export const postService = {
  async getById(postId: string): Promise<Post> {
    const response = await fetch(`${POST_URL}/${postId}`, {
      headers: defaultAuthHeaders()
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
    return response.json();
  },

  async getNearbyLosts(lat: number, lng: number, page: number): Promise<Post[]> {
    const url = `${POST_URL}/lost?lat=${lat}&lng=${lng}&page=${page}`;
    const response = await fetch(url);
    if (!response.ok) throw new HttpError(response.status, await response.json());
    return response.json();
  },

  async getByUsers(usernames: string[]): Promise<Post[]> {
    const response = await fetch(`${POST_URL}/ofUsers`, {
      method: 'POST',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ usernames })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
    return response.json();
  },

  async getByUsername(username: string): Promise<Post[]> {
    const response = await fetch(`${POST_URL}/ofUser/${username}`, {
      headers: defaultAuthHeaders()
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
    return response.json();
  },

  async getParticipatedBy(username: string): Promise<Post[]> {
    const response = await fetch(`${POST_URL}/ofUser/${username}/participated`, {
      headers: defaultAuthHeaders()
    })
    if (!response.ok) throw new HttpError(response.status, await response.json());
    return response.json();
  },

  async createBlogPost(payload: BlogPostPayload): Promise<Post> {
    const response = await fetch(`${POST_URL}/blog`, {
      method: 'POST',
      headers: defaultAuthHeaders(),
      body: JSON.stringify(payload)
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
    return response.json();
  },

  async createLostPost(payload: LostPostPayload): Promise<Post> {
    const response = await fetch(`${POST_URL}/lost`, {
      method: 'POST',
      headers: defaultAuthHeaders(),
      body: JSON.stringify(payload)
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
    return response.json();
  },

  async updateContent(postId: string, content: BlogPostPayload): Promise<Post> {
    const response = await fetch(`${POST_URL}/${postId}/content`, {
      method: 'PATCH',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ content })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
    return response.json();
  },

  async updateLostFoundStatus(postId: string, isResolved: boolean): Promise<void> {
    const response = await fetch(`${POST_URL}/${postId}/lostFoundStatus`, {
      method: 'PATCH',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ isResolved })
    });
    if (!response.ok) throw new HttpError(response.status, 'Failed to update lost found status');
  },

  async delete(postId: string): Promise<void> {
    const response = await fetch(`${POST_URL}`, {
      method: 'DELETE',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ postId })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
  },

  async getComments(postId: string): Promise<Comment[]> {
    const response = await fetch(`${POST_URL}/${postId}/comments`, {
      headers: defaultAuthHeaders()
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
    return (await response.json()).comments;
  },

  async addComment(postId: string, content: string): Promise<void> {
    const response = await fetch(`${POST_URL}/${postId}/comments`, {
      method: 'POST',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ content })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
  },

  async upsertInteraction(postId: string): Promise<void> {
    const response = await fetch(`${POST_URL}/${postId}/interactions`, {
      method: 'POST',
      headers: defaultAuthHeaders(),
      body: JSON.stringify({ type: 'like' })
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
  },

  async deleteInteraction(postId: string): Promise<void> {
    const response = await fetch(`${POST_URL}/${postId}/interactions`, {
      method: 'DELETE',
      headers: defaultAuthHeaders()
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
  },

  async participate(postId: string): Promise<void> {
    const response = await fetch(`${POST_URL}/${postId}/participation`, {
      method: 'POST',
      headers: defaultAuthHeaders()
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
  },

  async unparticipate(postId: string): Promise<void> {
    const response = await fetch(`${POST_URL}/${postId}/participation`, {
      method: 'DELETE',
      headers: defaultAuthHeaders()
    });
    if (!response.ok) throw new HttpError(response.status, await response.json());
  },
};