import React from 'react';
import { MapPin, Users, Calendar } from 'lucide-react';
import { Link } from 'react-router-dom';
import { Post } from '../../types/post';
import Avatar from '../common/Avatar';
import Card from '../common/Card';
import Button from '../common/Button';
import { formatDistanceToNow } from '../../utils/dateUtils';

interface LostPetCardProps {
  post: Post;
}

const LostPetCard: React.FC<LostPetCardProps> = ({ post }) => {
  const baseTagClassName = 'px-2 py-1 rounded-full text-xs font-medium ring-2 ring-white'
  return (
    <Card interactive className="h-full" bg={`${!post.isResolved ? 'bg-white-100' : 'bg-teal-50'}`}>
      <div className="relative">
        {post.medias?.length > 0 && (
          <img 
            src={post.medias[0].url} 
            alt="Pet"
            className="w-full h-48 object-cover"
          />
        )}
        <div className='absolute top-3 right-3 flex flex-row space-x-2'>
          <div className={`${baseTagClassName} ${!post.isResolved ? 'bg-error-100 text-error-700' : 'bg-neutral-100 text-neutral-700'}`}>
            {post.type === 'lost' ? 'Missing' : 'Found'}
          </div>
          { post.isResolved && 
            <div className={`${baseTagClassName} bg-success-100 text-success-700`}>Resolved</div> 
          }
        </div>
      </div>
      
      <div className="p-4">
        <h3 className="text-lg font-medium text-gray-900 mb-2">{post.content.split('\n')[0]}</h3>
        
        <div className="flex items-center space-x-2 text-sm text-gray-500 mb-3">
          <MapPin size={16} />
          <span>{post.area?.address}</span>
        </div>
        
        <div className="flex items-center justify-between text-sm text-gray-500 mb-4">
          <div className="flex items-center space-x-1">
            <Calendar size={16} />
            <span>{formatDistanceToNow(new Date(post.createdAt))}</span>
          </div>
          <div className="flex items-center space-x-1">
            <Users size={16} />
            <span>{post.participants?.length || 0} helpers</span>
          </div>
        </div>
        
        <div className="flex items-center justify-between">
          <div className="flex items-center space-x-2">
            <Avatar src={post.userAvatar} alt={post.displayName} size="sm" />
            <span className="text-sm text-gray-700">{post.displayName}</span>
          </div>
          
          <Link to={`/lost-pets/${post.id}`}>
            <Button variant="secondary" size="sm">
              View Details
            </Button>
          </Link>
        </div>
      </div>
    </Card>
  );
};

export default LostPetCard;