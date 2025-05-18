export interface GroupChat {
  id: number;
  name: string;
  isDirect: boolean;
  ownerName: string;
  members: string[];
  lastMessage: Message | null;
}

export interface Message {
  id: number;
  username: string;
  groupId: number;
  content: string;
  createdAt: string;
}