export interface Media {
  id: string;
  type: "image" | "video";
  url: string;
}

export interface Location {
  lat: number;
  lng: number;
  address: string;
}

export interface Interaction {
  type: "like" | "share";
  username: string;
}

export interface Post {
  id: string;
  type: "blog" | "lost" | "found";
  username: string;
  displayName: string;
  userAvatar: string;
  content: string;
  medias: Media[];
  createdAt: string;
  updatedAt: string;
  interactions: Interaction[];
  commentNum: number;

  // Optional: Lost Found Post
  lostAt?: string;
  area?: Location;
  lastSeen?: Location;
  contactInfo?: string;
  isResolved?: boolean;
  participants?: string[];
}

export interface Comment {
  username: string;
  displayName: string;
  avatar: string;
  content: string;
  createdAt: string;
}