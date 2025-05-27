import React, { useMemo } from 'react';
import { MapContainer, TileLayer, Marker } from 'react-leaflet';
import { DivIcon } from 'leaflet';
import 'leaflet/dist/leaflet.css';
import { Post } from '../../types/post';
import { calcDistance, formatDistance } from '../../utils/common';

interface LostPetMapProps {
  post: Post;
  userLocation?: { lat: number; lng: number };
}

const LostPetMap: React.FC<LostPetMapProps> = ({ post, userLocation }) => {
  if (post.type === 'blog') return;

  const lastSeenDistance = useMemo(() => {
    if (!userLocation || !post.lastSeen) return "0";
    return formatDistance(calcDistance(
      userLocation.lat, userLocation.lng,
      post.lastSeen.lat, post.lastSeen.lng
    ));
  }, [userLocation, post.lastSeen]);

  const areaDistance = useMemo(() => {
    if (!userLocation || !post.area) return "0";
    return formatDistance(calcDistance(
      userLocation.lat, userLocation.lng,
      post.area.lat, post.area.lng
    ));
  }, [userLocation, post.area]);

  const areaIcon = new DivIcon({
    className: '',
    html: `
      <div class="flex flex-col items-center translate-y-[-50%]">
        <div class="bg-white rounded-lg p-2 shadow-md text-center w-[110px]">
          <div class="text-xs font-bold text-gray-700">Area</div>
          ${ userLocation ? `<div class="text-xs text-gray-500">${areaDistance} away</div>` : ''}
        </div>
        <div class="w-0 h-0 border-l-[8px] border-l-transparent border-r-[8px] border-r-transparent border-t-[10px] border-t-white -mt-[1px]"></div>
      </div>
    `,
  });

  const lastSeenIcon = new DivIcon({
    className: '',
    html: `
      <div class="flex flex-col items-center translate-y-[-50%]">
        <div class="bg-white rounded-lg p-2 shadow-md text-center w-[110px]">
          <div class="text-xs font-bold text-gray-700">Last Seen</div>
          ${ userLocation ? `<div class="text-xs text-gray-500">${lastSeenDistance} away</div>` : ''}
        </div>
        <div class="w-0 h-0 border-l-[8px] border-l-transparent border-r-[8px] border-r-transparent border-t-[10px] border-t-white -mt-[1px]"></div>
      </div>
    `,
  });

  const userIcon = new DivIcon({
    className: 'bg-primary-600 rounded-full shadow-md',
    iconSize: [22, 22],
  });

  return (
    <div className="h-[400px] mb-4">
      <MapContainer
        center={[post.lastSeen?.lat || 0, post.lastSeen?.lng || 0]}
        zoom={13}
        style={{ height: '100%', width: '100%' }}
      >
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        />
        {post.lastSeen && (
          <Marker position={[post.lastSeen.lat, post.lastSeen.lng]} icon={lastSeenIcon} />
        )}
        {post.area && (
          <Marker position={[post.area.lat, post.area.lng]} icon={areaIcon} />
        )}
        {userLocation && (
          <Marker position={[userLocation.lat, userLocation.lng]} icon={userIcon} />
        )}
      </MapContainer>
    </div>
  );
};

export default LostPetMap;