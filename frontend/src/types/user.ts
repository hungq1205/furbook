export interface User {
  username: string;
  displayName: string;
  avatar: string;
  bio: string;
  friendNum: number;
  groupid?: number;
}

export type Friendship = 'none' | 'sent' | 'received' | 'friend';