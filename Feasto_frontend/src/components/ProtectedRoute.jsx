import { Navigate } from "react-router-dom";
import { getUserFromToken } from "../utils/Auth";

export function AdminRoute({ children }) {
  const user = getUserFromToken();
  if (!user) return <Navigate to="/" />; // not logged in


  const role = user.user_role
  if (role !== "admin") {
    return <Navigate to="/" />;
  }

  return children;
}

export function ProtectedRoute({ children, role1 , role2 }) {
  const user = getUserFromToken();
  
  if (!user) return <Navigate to="/" />; // not logged in
  if ((role1 || role2) && user.user_role !== role1 && user.user_role !== role2) return <Navigate to="/unauthorized" />; // role mismatch
  
  return children;
}
