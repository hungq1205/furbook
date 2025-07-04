import React, { useMemo } from 'react';
import { MapPin, Users, Calendar } from 'lucide-react';
import { Link } from 'react-router-dom';
import { Post } from '../../types/post';
import Avatar from '../common/Avatar';
import Card from '../common/Card';
import Button from '../common/Button';
import { calcDistance, formatDistance, formatDistanceToNow } from '../../utils/common';

interface LostPetCardProps {
  post: Post;
  userLocation: {lat: number, lng: number} | undefined;
}

export const getTagColor = (type: 'lost' | 'found' | 'blog', isResolved?: boolean): string => {
  if (isResolved)
    return 'bg-neutral-100 text-neutral-700';
  return type === 'lost' 
    ? 'bg-error-100 text-error-700' 
    : 'bg-orange-100 text-orange-700';
};

const LostPetCard: React.FC<LostPetCardProps> = ({ post, userLocation }) => {
  const baseTagClassName = 'px-2 py-1 rounded-full text-xs font-medium ring-2 ring-white';
  const tagColor = getTagColor(post.type, post.isResolved);

  const distance = useMemo(() => {
    if (!userLocation || !post.lastSeen) return "0";
    return formatDistance(calcDistance(
      userLocation.lat, userLocation.lng,
      post.lastSeen.lat, post.lastSeen.lng
    ));
  }, [userLocation, post.lastSeen]);

  return (
    <Card interactive className="h-full flex flex-col" bg={`${!post.isResolved ? 'bg-white-100' : 'bg-teal-50'}`}>
      <div className="relative">
        {post.medias?.length > 0 && (
          <Link to={`/lost-pets/${post.id}`}>
            <img 
              src={post.medias[0].url} 
              alt="Pet"
              className="w-full h-48 object-cover"
            />
          </Link>
        )}
        <div className='absolute top-3 right-3 flex flex-row space-x-2'>
          { !post.isResolved && userLocation && 
            <div className={`${baseTagClassName} bg-sky-100 text-sky-700`}>{distance} away</div>
          }
          <div className={`${baseTagClassName} ${tagColor}`}>
            {post.type === 'lost' ? 'Missing' : 'Found'}
          </div>
          { post.isResolved &&
            <div className={`${baseTagClassName} bg-success-100 text-success-700`}>Resolved</div>
          }
        </div>
      </div>
      
      <div className="p-4 flex flex-col grow">
        <Link className="grow" to={`/lost-pets/${post.id}`}>
          <h3 className="text-lg font-medium text-gray-900 mb-2">{post.content.split('\n')[0]}</h3>
        </Link>

        <div className="flex items-center space-x-2 text-sm text-gray-500 mb-3">
          <MapPin size={16} />
          <span>{post.lastSeen?.address}</span>
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
            <Link to={`/profile/${post.username}`}>
              <Avatar src={post.userAvatar} alt={post.displayName} size="sm" />
            </Link>
            <Link to={`/profile/${post.username}`}>
              <span className="text-sm text-gray-700">{post.displayName}</span>
            </Link>
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