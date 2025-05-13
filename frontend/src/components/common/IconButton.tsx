import React from 'react';
import { motion } from 'framer-motion';

interface IconButtonProps {
  icon: React.ReactNode;
  onClick?: () => void;
  size?: 'sm' | 'md' | 'lg';
  variant?: 'primary' | 'secondary' | 'outline' | 'ghost';
  label: string; 
  badge?: number;
}

const IconButton: React.FC<IconButtonProps> = ({
  icon,
  onClick,
  size = 'md',
  variant = 'ghost',
  label,
  badge,
}) => {
  const sizeClasses = {
    sm: 'p-1.5',
    md: 'p-2',
    lg: 'p-3',
  };

  const variantClasses = {
    primary: 'bg-primary-600 text-white hover:bg-primary-700',
    secondary: 'bg-secondary-600 text-white hover:bg-secondary-700',
    outline: 'border border-gray-300 bg-transparent text-gray-700 hover:bg-gray-50',
    ghost: 'bg-transparent text-gray-700 hover:bg-gray-100',
  };

  return (
    <motion.button
      aria-label={label}
      onClick={onClick}
      className={`relative rounded-full ${sizeClasses[size]} ${variantClasses[variant]} transition-colors focus:outline-none`}
      whileTap={{ scale: 0.95 }}
    >
      {icon}
      {badge !== undefined && badge > 0 && (
        <span className="absolute -top-1 -right-1 flex h-5 w-5 items-center justify-center rounded-full bg-error-500 text-xs font-medium text-white">
          {badge > 9 ? '9+' : badge}
        </span>
      )}
    </motion.button>
  );
};

export default IconButton;