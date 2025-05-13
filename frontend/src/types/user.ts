export interface User {
  username: string;
  displayName: string;
  avatar: string;
  bio: string;
  friendNum: number;
}

export interface GroupChat {
  id: number;
  name: string;
  isDirect: boolean;
  ownerName: string;
  members: string[];
}

export interface Message {
  id: number;
  username: string;
  groupId: number;
  content: string;
  createdAt: string;
}