import React, { useState } from 'react';
import { ChevronLeft, ChevronRight } from 'lucide-react';
import { Media } from '../../types/post';

interface MediaGalleryProps {
  media: Media[];
  className?: string;
  crop?: boolean;
}

const MediaGallery: React.FC<MediaGalleryProps> = ({ media, className = '', crop = true }) => {
  const [currentIndex, setCurrentIndex] = useState(0);

  if (!media || media.length === 0) {
    return null;
  }

  const handlePrevious = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.stopPropagation();
    setCurrentIndex((prev) => (prev === 0 ? media.length - 1 : prev - 1));
  };

  const handleNext = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.stopPropagation();
    setCurrentIndex((prev) => (prev === media.length - 1 ? 0 : prev + 1));
  };

  const renderMedia = (item: Media) => {
    if (item.type === 'image') {
      return (
        <img 
          src={item.url} 
          alt="Post media"
          className="w-full h-full object-cover"
        />
      );
    } else if (item.type === 'video') {
      return (
        <video 
          src={item.url}
          className="w-full h-full object-cover"
          controls
        />
      );
    }
    return null;
  };

  return (
    <div className={`relative overflow-hidden rounded-md ${className}`}>
      <div className="aspect-video bg-gray-100 relative">
        <div className={crop ? "w-full h-full" : "w-full"}>
          {renderMedia(media[currentIndex])}
        </div>
      </div>
      
      {media.length > 1 && (
        <>
          <button 
            onClick={handlePrevious}
            className="absolute left-2 top-1/2 -translate-y-1/2 bg-black/30 text-white p-1 rounded-full hover:bg-black/50 transition-colors z-10"
            aria-label="Previous media"
          >
            <ChevronLeft size={20} />
          </button>
          <button 
            onClick={handleNext}
            className="absolute right-2 top-1/2 -translate-y-1/2 bg-black/30 text-white p-1 rounded-full hover:bg-black/50 transition-colors z-10"
            aria-label="Next media"
          >
            <ChevronRight size={20} />
          </button>
          <div className="absolute bottom-2 left-1/2 -translate-x-1/2 flex space-x-1">
            {media.map((_, idx) => (
              <button
                key={idx}
                className={`w-2 h-2 rounded-full ${
                  idx === currentIndex ? 'bg-white' : 'bg-white/50'
                }`}
                onClick={() => setCurrentIndex(idx)}
                aria-label={`Go to media ${idx + 1}`}
              />
            ))}
          </div>
        </>
      )}
    </div>
  );
};

export default MediaGallery;