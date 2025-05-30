import React from 'react';
import { motion } from 'framer-motion';

interface AvatarProps {
  src: string;
  alt: string;
  className?: string;
  size?: 'sm' | 'md' | 'lg' | 'xl';
  isOnline?: boolean;
  onClick?: () => void;
}

const Avatar: React.FC<AvatarProps> = ({ 
  src, 
  alt, 
  className,
  size = 'md',
  isOnline,
  onClick,
}) => {
  const sizeClasses = {
    sm: 'w-8 h-8',
    md: 'w-10 h-10',
    lg: 'w-14 h-14',
    xl: 'w-24 h-24'
  };

  return (
    <div className={`${sizeClasses[size]} relative ${className}`} onClick={onClick}>
      <motion.img 
        src={src || 'https://phanmemmkt.vn/wp-content/uploads/2024/09/Hinh-anh-dai-dien-mac-dinh-Facebook.jpg'} 
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