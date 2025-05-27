import React from 'react';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import { DivIcon } from 'leaflet';
import 'leaflet/dist/leaflet.css';
import { Link } from 'react-router-dom';
import { Post } from '../../types/post';
import { calcDistance, formatDistance } from '../../utils/common';

interface LostPetsMapProps {
  posts: Post[];
  userLocation?: { lat: number; lng: number };
}

const LostPetsMap: React.FC<LostPetsMapProps> = ({ posts, userLocation }) => {
  const validPosts = posts.filter(post => post.lastSeen?.lat && post.lastSeen?.lng && !post.isResolved);
  const createMarkerImage = (post: Post) => {
    const imageUrl = post.medias.find(m => m.type === 'image')?.url || 'https://png.pngtree.com/png-vector/20191005/ourmid/pngtree-animal-paw-print-icon-png-image_1794752.jpg';
    const color = post.type === 'lost' ? 'error-100' : 'orange-100'

    return new DivIcon({
      className: '',
      html: `
        <div class="flex flex-col items-center translate-y-[-50%]">
          <div class="bg-${color} rounded-2xl p-2 shadow-md text-center w-[70px]">
            <div class="text-xs font-bold mb-1 text-gray-700">${post.type === 'lost' ? 'Missing' : 'Found'}</div>
              <img src="${imageUrl}" class="object-cover rounded-lg" style="width: 100%; height: 100%; max-height: 40px !important;" />
            </div>
          <div class="w-0 h-0 border-l-[8px] border-l-transparent border-r-[8px] border-r-transparent border-t-[10px] border-t-${color} -mt-[1px] shadow-md"></div>
        </div>
      `,
    });
  };

  const userIcon = new DivIcon({
    className: 'bg-primary-600 rounded-full shadow-md',
    iconSize: [22, 22],
  });

  return (
    <div className="h-[400px] rounded-lg overflow-hidden mb-6">
      <MapContainer
        center={userLocation ? [userLocation.lat, userLocation.lng] : [0, 0]}
        zoom={13}
        style={{ height: '100%', width: '100%' }}
      >
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        />
        
        {validPosts.map((post) => (
          <Marker
            key={post.id}
            position={[post.lastSeen!.lat, post.lastSeen!.lng]}
            icon={createMarkerImage(post)}
          >
            <Popup>
              <div className="p-2">
                <h3 className="font-medium text-sm mb-1">{post.content.split('\n')[0]}</h3>
                <p className="text-xs text-gray-500 mb-2">{post.lastSeen?.address}</p>
                <Link
                  to={`/lost-pets/${post.id}`}
                  className="text-xs text-primary-600 hover:underline"
                >
                  View Details
                </Link>
                { !post.isResolved && userLocation && post.lastSeen &&
                  <div className="absolute right-6 bottom-5 px-3 py-1 rounded-full text-sm font-medium bg-sky-100 text-sky-700">
                  {formatDistance(calcDistance(userLocation.lat, userLocation.lng, post.lastSeen?.lat, post.lastSeen?.lng))} away
                  </div>
                }
              </div>
            </Popup>
          </Marker>
        ))}
        {userLocation && (
          <Marker position={[userLocation.lat, userLocation.lng]} icon={userIcon} />
        )}
      </MapContainer>
    </div>
  );
};

export default LostPetsMap;