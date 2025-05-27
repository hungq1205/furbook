import React, { useState } from 'react';
import { MapContainer, TileLayer, Marker, useMapEvents } from 'react-leaflet';
import { DivIcon, LeafletMouseEvent } from 'leaflet';
import 'leaflet/dist/leaflet.css';
import Button from '../common/Button';
import { Location } from '../../types/post';

interface LocationPickerProps {
  onSelected: (location: Location) => void;
  onClose: () => void;
  userLocation?: { lat: number; lng: number };
}

const MapEvents = ({ onLocationSelect }: { onLocationSelect: (lat: number, lng: number) => void }) => {
  useMapEvents({
    click: (e: LeafletMouseEvent) => {
      onLocationSelect(e.latlng.lat, e.latlng.lng);
    },
  });
  return null;
};

const LocationPicker: React.FC<LocationPickerProps> = ({ onSelected, onClose, userLocation }) => {
  const [marker, setMarker] = useState<{ lat: number; lng: number } | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [address, setAddress] = useState<string | null>(null);

  const fetchAddress = async (lat: number, lng: number) => {
    try {
      const response = await fetch(`https://nominatim.openstreetmap.org/reverse?lat=${lat}&lon=${lng}&format=json`);
      if (!response.ok) throw new Error('Failed to fetch address');
      const data = await response.json();
      setAddress(data.display_name || `${lat}, ${lng}`);
    } catch (error) {
      console.error('Error fetching address:', error);
      setAddress(null);
    } finally {
      setIsLoading(false);
    }
  };

  const handleLocationSelect = async (lat: number, lng: number) => {
    setMarker({ lat, lng });
    setIsLoading(true);
    fetchAddress(lat, lng);
  };

  const markerIcon = new DivIcon({
    className: 'bg-orange-600 rounded-full shadow-md',
    iconSize: [16, 16],
  });
  
  const userIcon = new DivIcon({
    className: 'bg-primary-600 rounded-full shadow-md',
    iconSize: [22, 22],
  });

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-white rounded-lg p-4 w-full max-w-3xl">
        <div className="mb-4">
          <h2 className="text-xl font-semibold">
          { (isLoading ? 'Loading...' : address) || 'Select Location' }
          </h2>
          { marker ? 
            <>
              <p className="text-sm text-gray-500">{`Longitude: ${marker?.lng.toFixed(3)}`}</p>
              <p className="text-sm text-gray-500">{`Latitude: ${marker?.lat.toFixed(3)}`}</p>
            </> :
            <p className="text-sm text-gray-500">Click on the map to set the location</p> 
          }
        </div>
        
        <div className="h-[400px] mb-4">
          <MapContainer
            center={[userLocation?.lat || 0, userLocation?.lng || 0]}
            zoom={13}
            style={{ height: '100%', width: '100%' }}
          >
            <TileLayer
              url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
              attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
            />
            <MapEvents onLocationSelect={handleLocationSelect} />
            {userLocation && (
              <Marker position={[userLocation.lat, userLocation.lng]} icon={userIcon} />
            )}
            {marker && (
              <Marker position={[marker.lat, marker.lng]} icon={markerIcon} />
            )}
          </MapContainer>
        </div>
        
        <div className="flex justify-end space-x-3">
          <Button variant="ghost" onClick={onClose}>Cancel</Button>
          <Button 
            variant="primary" 
            onClick={() => {
              onSelected({ lat: marker!.lat, lng: marker!.lng, address: address || '' });
              onClose();
            }}
            disabled={!marker || isLoading}
          >
            {isLoading ? 'Loading...' : 'Confirm Location'}
          </Button>
        </div>
      </div>
    </div>
  );
};

export default LocationPicker;