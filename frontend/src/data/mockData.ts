import { Post } from '../types/post';
import { User, GroupChat, Message } from '../types/user';

export const currentUser: User = {
  username: 'janecooper',
  displayName: 'Jane Cooper',
  avatar: 'https://images.pexels.com/photos/733872/pexels-photo-733872.jpeg?auto=compress&cs=tinysrgb&w=150',
  bio: 'Dog lover, adventure seeker, and full-time pet parent. I love sharing my furry companions with the world!',
  friendNum: 128
};

export const users: User[] = [
  {
    username: 'alexmorgan',
    displayName: 'Alex Morgan',
    avatar: 'https://images.pexels.com/photos/220453/pexels-photo-220453.jpeg?auto=compress&cs=tinysrgb&w=150',
    bio: 'Cat enthusiast and photographer',
    friendNum: 86
  },
  {
    username: 'sarahjohnson',
    displayName: 'Sarah Johnson',
    avatar: 'https://images.pexels.com/photos/415829/pexels-photo-415829.jpeg?auto=compress&cs=tinysrgb&w=150',
    bio: 'Proud mom of 2 golden retrievers',
    friendNum: 102
  },
  {
    username: 'michaelchen',
    displayName: 'Michael Chen',
    avatar: 'https://images.pexels.com/photos/2379004/pexels-photo-2379004.jpeg?auto=compress&cs=tinysrgb&w=150',
    bio: 'Exotic pet enthusiast',
    friendNum: 93
  },
  {
    username: 'emilydavis',
    displayName: 'Emily Davis',
    avatar: 'https://images.pexels.com/photos/774909/pexels-photo-774909.jpeg?auto=compress&cs=tinysrgb&w=150',
    bio: 'Rescue advocate and volunteer',
    friendNum: 145
  },
  {
    username: 'carlosrodriguez',
    displayName: 'Carlos Rodriguez',
    avatar: 'https://images.pexels.com/photos/1300402/pexels-photo-1300402.jpeg?auto=compress&cs=tinysrgb&w=150',
    bio: 'Professional dog trainer',
    friendNum: 214
  }
];

export const friends: User[] = [
  {
    username: 'alexmorgan',
    displayName: 'Alex Morgan',
    avatar: 'https://images.pexels.com/photos/220453/pexels-photo-220453.jpeg?auto=compress&cs=tinysrgb&w=150',
    bio: 'Cat enthusiast and photographer',
    friendNum: 86
  },
  {
    username: 'sarahjohnson',
    displayName: 'Sarah Johnson',
    avatar: 'https://images.pexels.com/photos/415829/pexels-photo-415829.jpeg?auto=compress&cs=tinysrgb&w=150',
    bio: 'Proud mom of 2 golden retrievers',
    friendNum: 102
  },
  {
    username: 'michaelchen',
    displayName: 'Michael Chen',
    avatar: 'https://images.pexels.com/photos/2379004/pexels-photo-2379004.jpeg?auto=compress&cs=tinysrgb&w=150',
    bio: 'Exotic pet enthusiast',
    friendNum: 93
  }
];

export const groupChats: GroupChat[] = [
  {
    id: 1,
    name: 'Pet Lovers Club',
    isDirect: false,
    ownerName: 'janecooper',
    members: ['janecooper', 'alexmorgan', 'sarahjohnson', 'michaelchen']
  },
  {
    id: 2,
    name: 'Dog Training Tips',
    isDirect: false,
    ownerName: 'carlosrodriguez',
    members: ['carlosrodriguez', 'janecooper', 'sarahjohnson']
  },
  {
    id: 3,
    name: 'Alex Morgan',
    isDirect: true,
    ownerName: 'janecooper',
    members: ['janecooper', 'alexmorgan']
  },
  {
    id: 4,
    name: 'Sarah Johnson',
    isDirect: true,
    ownerName: 'janecooper',
    members: ['janecooper', 'sarahjohnson']
  }
];

export const messages: Message[] = [
  {
    id: 1,
    username: 'janecooper',
    groupId: 1,
    content: 'Hey everyone! How are your pets doing today?',
    createdAt: '2024-03-15T10:30:00Z'
  },
  {
    id: 2,
    username: 'alexmorgan',
    groupId: 1,
    content: 'My cat just learned a new trick!',
    createdAt: '2024-03-15T10:32:00Z'
  },
  {
    id: 3,
    username: 'sarahjohnson',
    groupId: 1,
    content: 'That\'s amazing! What trick?',
    createdAt: '2024-03-15T10:33:00Z'
  },
  {
    id: 4,
    username: 'carlosrodriguez',
    groupId: 2,
    content: 'Here\'s a tip: Always reward good behavior immediately.',
    createdAt: '2024-03-15T11:00:00Z'
  },
  {
    id: 5,
    username: 'janecooper',
    groupId: 2,
    content: 'That\'s really helpful, thanks Carlos!',
    createdAt: '2024-03-15T11:02:00Z'
  },
  {
    id: 6,
    username: 'janecooper',
    groupId: 3,
    content: 'Hey Alex, want to meet up at the pet park?',
    createdAt: '2024-03-15T12:00:00Z'
  },
  {
    id: 7,
    username: 'alexmorgan',
    groupId: 3,
    content: 'Sure! How about 3pm?',
    createdAt: '2024-03-15T12:05:00Z'
  }
];

export const posts: Post[] = [
  {
    id: '1',
    type: 'blog',
    username: 'alexmorgan',
    displayName: 'Alex Morgan',
    userAvatar: 'https://images.pexels.com/photos/220453/pexels-photo-220453.jpeg?auto=compress&cs=tinysrgb&w=150',
    content: 'My cat Whiskers is enjoying the sunshine today! Who else has pets that love sunbathing?',
    medias: [
      {
        id: '1',
        type: 'image',
        url: 'https://images.pexels.com/photos/2061057/pexels-photo-2061057.jpeg?auto=compress&cs=tinysrgb&w=600'
      }
    ],
    createdAt: '2024-03-15T10:30:00Z',
    updatedAt: '2024-03-15T10:30:00Z',
    interactions: [
      { type: 'like', username: 'janecooper' },
      { type: 'like', username: 'sarahjohnson' }
    ],
    commentNum: 7
  },
  {
    id: '2',
    type: 'lost',
    username: 'emilydavis',
    displayName: 'Emily Davis',
    userAvatar: 'https://images.pexels.com/photos/774909/pexels-photo-774909.jpeg?auto=compress&cs=tinysrgb&w=150',
    content: 'My tabby cat Felix has been missing since yesterday evening. He\'s orange with white paws, very friendly, and responds to his name.',
    medias: [
      {
        id: '2',
        type: 'image',
        url: 'https://images.pexels.com/photos/2071873/pexels-photo-2071873.jpeg?auto=compress&cs=tinysrgb&w=600'
      }
    ],
    createdAt: '2024-03-14T09:15:00Z',
    updatedAt: '2024-03-14T14:20:00Z',
    interactions: [],
    commentNum: 15,
    lostAt: '2024-03-13T20:00:00Z',
    area: {
      lat: 40.7128,
      lng: -74.0060,
      address: 'Downtown West'
    },
    lastSeen: {
      lat: 40.7129,
      lng: -74.0061,
      address: 'Oak Street Park'
    },
    contactInfo: '555-123-4567',
    isResolved: false,
    participants: ['janecooper', 'alexmorgan', 'sarahjohnson']
  },
  {
    id: '3',
    type: 'blog',
    username: 'janecooper',
    displayName: 'Jane Cooper',
    userAvatar: 'https://images.pexels.com/photos/733872/pexels-photo-733872.jpeg?auto=compress&cs=tinysrgb&w=150',
    content: 'First day at puppy training class! Max is doing so well with basic commands. 🐕',
    medias: [
      {
        id: '3',
        type: 'image',
        url: 'https://images.pexels.com/photos/1805164/pexels-photo-1805164.jpeg?auto=compress&cs=tinysrgb&w=600'
      }
    ],
    createdAt: '2024-03-15T15:45:00Z',
    updatedAt: '2024-03-15T15:45:00Z',
    interactions: [
      { type: 'like', username: 'alexmorgan' },
      { type: 'like', username: 'michaelchen' },
      { type: 'share', username: 'sarahjohnson' }
    ],
    commentNum: 4
  },
  {
    id: '4',
    type: 'found',
    username: 'michaelchen',
    displayName: 'Michael Chen',
    userAvatar: 'https://images.pexels.com/photos/2379004/pexels-photo-2379004.jpeg?auto=compress&cs=tinysrgb&w=150',
    content: 'Found this sweet golden retriever near Central Park. No collar but very well-trained. Please share to help find the owner!',
    medias: [
      {
        id: '4',
        type: 'image',
        url: 'https://images.pexels.com/photos/1490908/pexels-photo-1490908.jpeg?auto=compress&cs=tinysrgb&w=600'
      }
    ],
    createdAt: '2024-03-16T09:20:00Z',
    updatedAt: '2024-03-16T09:20:00Z',
    interactions: [
      { type: 'share', username: 'janecooper' },
      { type: 'share', username: 'emilydavis' }
    ],
    commentNum: 8,
    area: {
      lat: 40.7829,
      lng: -73.9654,
      address: 'Central Park East'
    },
    lastSeen: {
      lat: 40.7831,
      lng: -73.9652,
      address: 'East 72nd Street Entrance'
    },
    contactInfo: '555-987-6543',
    isResolved: false,
    participants: ['janecooper', 'emilydavis', 'sarahjohnson', 'alexmorgan']
  },
  {
    id: '5',
    type: 'blog',
    username: 'sarahjohnson',
    displayName: 'Sarah Johnson',
    userAvatar: 'https://images.pexels.com/photos/415829/pexels-photo-415829.jpeg?auto=compress&cs=tinysrgb&w=150',
    content: 'Sunday funday at the dog park! Luna and Max made so many new friends today. 🐾',
    medias: [
      {
        id: '5',
        type: 'image',
        url: 'https://images.pexels.com/photos/1108099/pexels-photo-1108099.jpeg?auto=compress&cs=tinysrgb&w=600'
      },
      {
        id: '6',
        type: 'image',
        url: 'https://images.pexels.com/photos/2607544/pexels-photo-2607544.jpeg?auto=compress&cs=tinysrgb&w=600'
      }
    ],
    createdAt: '2024-03-17T16:30:00Z',
    updatedAt: '2024-03-17T16:30:00Z',
    interactions: [
      { type: 'like', username: 'janecooper' },
      { type: 'like', username: 'alexmorgan' },
      { type: 'like', username: 'emilydavis' },
      { type: 'share', username: 'michaelchen' }
    ],
    commentNum: 12
  }
];