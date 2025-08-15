import { jwtDecode } from "jwt-decode";

function getCookie(name) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop().split(";").shift();
  return null;
}

export function getUserFromToken() {
  const token = getCookie("auth_token");
  if (!token) return null;

  try {
    const decoded = jwtDecode(token);

    // Check if expired
    if (decoded.exp * 1000 < Date.now()) {
      clearCookie("auth_token"); // uses Max-Age=0
      return null;
    }

    return decoded;
  } catch {
    return null;
  }
}

export const clearCookie = (name) => {
  document.cookie = `${name}=; Max-Age=0; path=/`;
};


