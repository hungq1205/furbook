import { Post, Comment } from '../types/post';
import { User, GroupChat, Message } from '../types/user';
import { authService } from './auth';

const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:3000/api';

const getHeaders = () => ({
  'Content-Type': 'application/json',
  ...authService.getAuthHeaders()
});

// User API
export const userApi = {
  async getByUsername(username: string): Promise<User> {
    const response = await fetch(`${BASE_URL}/users/${username}`, {
      headers: getHeaders()
    });
    if (!response.ok) throw new Error('Failed to fetch user');
    return response.json();
  },

  async getMultiple(usernames: string[]): Promise<User[]> {
    const queryString = usernames.join(',');
    const response = await fetch(`${BASE_URL}/users?usernames=${queryString}`, {
      headers: getHeaders()
    });
    if (!response.ok) throw new Error('Failed to fetch users');
    return response.json();
  }
};

// Chat API
export const chatApi = {
  async getGroups(): Promise<GroupChat[]> {
    const response = await fetch(`${BASE_URL}/chats`, {
      headers: getHeaders()
    });
    if (!response.ok) throw new Error('Failed to fetch chats');
    return response.json();
  },

  async getMessages(groupId: number): Promise<Message[]> {
    const response = await fetch(`${BASE_URL}/chats/${groupId}/messages`, {
      headers: getHeaders()
    });
    if (!response.ok) throw new Error('Failed to fetch messages');
    return response.json();
  },

  async sendMessage(groupId: number, content: string): Promise<Message> {
    const response = await fetch(`${BASE_URL}/chats/${groupId}/messages`, {
      method: 'POST',
      headers: getHeaders(),
      body: JSON.stringify({ content })
    });
    if (!response.ok) throw new Error('Failed to send message');
    return response.json();
  },

  async createGroup(name: string, members: string[]): Promise<GroupChat> {
    const response = await fetch(`${BASE_URL}/chats`, {
      method: 'POST',
      headers: getHeaders(),
      body: JSON.stringify({ name, members })
    });
    if (!response.ok) throw new Error('Failed to create group');
    return response.json();
  }
};

// Post API
export const postApi = {
  async getAll(): Promise<Post[]> {
    const response = await fetch(`${BASE_URL}/posts`, {
      headers: getHeaders()
    });
    if (!response.ok) throw new Error('Failed to fetch posts');
    return response.json();
  },

  async getById(postId: string): Promise<Post> {
    const response = await fetch(`${BASE_URL}/posts/${postId}`, {
      headers: getHeaders()
    });
    if (!response.ok) throw new Error('Failed to fetch post');
    return response.json();
  },

  async create(post: Omit<Post, 'id' | 'createdAt' | 'updatedAt'>): Promise<Post> {
    const response = await fetch(`${BASE_URL}/posts`, {
      method: 'POST',
      headers: getHeaders(),
      body: JSON.stringify(post)
    });
    if (!response.ok) throw new Error('Failed to create post');
    return response.json();
  },

  async update(postId: string, updates: Partial<Post>): Promise<Post> {
    const response = await fetch(`${BASE_URL}/posts/${postId}`, {
      method: 'PATCH',
      headers: getHeaders(),
      body: JSON.stringify(updates)
    });
    if (!response.ok) throw new Error('Failed to update post');
    return response.json();
  },

  async delete(postId: string): Promise<void> {
    const response = await fetch(`${BASE_URL}/posts/${postId}`, {
      method: 'DELETE',
      headers: getHeaders()
    });
    if (!response.ok) throw new Error('Failed to delete post');
  },

  async getComments(postId: string): Promise<Comment[]> {
    const response = await fetch(`${BASE_URL}/posts/${postId}/comments`, {
      headers: getHeaders()
    });
    if (!response.ok) throw new Error('Failed to fetch comments');
    return response.json();
  },

  async addComment(postId: string, content: string): Promise<Comment> {
    const response = await fetch(`${BASE_URL}/posts/${postId}/comments`, {
      method: 'POST',
      headers: getHeaders(),
      body: JSON.stringify({ content })
    });
    if (!response.ok) throw new Error('Failed to add comment');
    return response.json();
  }
};