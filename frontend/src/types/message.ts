export interface GroupChat {
  id: number;
  name: string;
  is_direct: boolean;
  owner_name: string;
  members: string[];
  last_message: Message | null;

  avatar?: string;
}

export interface Message {
  id: number;
  username: string;
  group_id: number;
  content: string;
  created_at: string;
}