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
    // Check expiration
    if (decoded.exp * 1000 < Date.now()) {
      // Optional: Clear the cookie by setting past expiry date
      document.cookie = "token=; path=/; expires=Thu, 01 Jan 1970 00:00:00 UTC;";
      return null;
    }
    return decoded;
  } catch (err) {
    return null;
  }
}

