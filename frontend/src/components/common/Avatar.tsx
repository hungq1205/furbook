import React from 'react';
import { motion } from 'framer-motion';

interface AvatarProps {
  src: string;
  alt: string;
  size?: 'sm' | 'md' | 'lg' | 'xl';
  isOnline?: boolean;
}

const Avatar: React.FC<AvatarProps> = ({ 
  src, 
  alt, 
  size = 'md',
  isOnline
}) => {
  const sizeClasses = {
    sm: 'w-8 h-8',
    md: 'w-10 h-10',
    lg: 'w-14 h-14',
    xl: 'w-24 h-24'
  };

  return (
    <div className={`${sizeClasses[size]} relative`}>
      <motion.img 
        src={src} 
        alt={alt}
        className={`${sizeClasses[size]} rounded-full object-cover border-2 border-white`}
        whileHover={{ scale: 1.05 }}
        transition={{ duration: 0.2 }}
      />
      {isOnline && (
        <span className="absolute bottom-0 right-0 block h-3 w-3 rounded-full bg-success-500 ring-2 ring-white" />
      )}
    </div>
  );
};

export default Avatar;