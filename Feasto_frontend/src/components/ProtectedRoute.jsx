import { Navigate } from "react-router-dom";
import { getUserFromToken } from "../utils/auth";

export function AdminRoute({ children }) {
  const user = getUserFromToken();
  if (!user) return <Navigate to="/" />; // not logged in


  const role = user.user_role
  if (role !== "admin") {
    return <Navigate to="/" />;
  }

  return children;
}

export function ProtectedRoute({ children, role }) {
  const user = getUserFromToken();
  
  if (!user) return <Navigate to="/" />; // not logged in
  if (role && user.user_role !== role) return <Navigate to="/unauthorized" />; // role mismatch
  
  return children;
}