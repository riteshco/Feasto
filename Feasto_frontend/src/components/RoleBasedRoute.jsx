import React from "react";
import { Navigate } from "react-router-dom";

import { ChefHome } from "@/pages/home/ChefHome";
import { CustomerHome } from "@/pages/home/CustomerHome";
import { getUserFromToken } from "@/utils/Auth";
import { ProtectedRoute } from "./ProtectedRoute";

export function RoleBasedRoute() {
  const user = getUserFromToken()
  if (!user) {
    return <Navigate to="/" replace />;
  }
    const role = user.user_role

    switch (role) {
      case "chef":
        return <ChefHome />;
      case "customer":
        return (
            <CustomerHome />
      )
      default:
        return <h1>Unauthorized</h1>;
    }
}
