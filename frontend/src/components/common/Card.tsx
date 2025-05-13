import React from 'react';
import { motion } from 'framer-motion';

interface CardProps {
  children: React.ReactNode;
  className?: string;
  onClick?: () => void;
  interactive?: boolean;
}

const Card: React.FC<CardProps> = ({ 
  children, 
  className = '', 
  onClick,
  interactive = false,
}) => {
  return (
    <motion.div
      className={`bg-white rounded-lg shadow-sm overflow-hidden ${className}`}
      onClick={onClick}
      whileHover={interactive ? { y: -4, boxShadow: '0 10px 15px -3px rgba(0, 0, 0, 0.1)' } : {}}
      transition={{ duration: 0.2 }}
    >
      {children}
    </motion.div>
  );
};

export default Card;