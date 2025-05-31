import { useEffect, useState } from "react";
import { useAuth } from "../../services/authService";
import wsService from "../../services/webSocketService";
import { Navigate } from "react-router-dom";

const ProtectedRoute = ({ children }: { children: JSX.Element }) => {
  const [authChecked, setAuthChecked] = useState(false);
  const authService = useAuth();

  useEffect(() => {
    authService.refresh()
      .then(() => {
        setAuthChecked(true);
      })
      .catch((error) => {
        console.error('Error checking authentication:', error);
        authService.logout();
        setAuthChecked(false);
      });
  }, []);

  useEffect(() => {
    if (!authService.isAuthenticated) return;
    wsService.connect(authService.token!);
    return () => wsService.disconnect();
  }, [authService.currentUser?.username]);

  if (!authChecked)
    return <div>Loading...</div>;

  return authService.isAuthenticated ? children : <Navigate to="/login" replace />;
};

export default ProtectedRoute;