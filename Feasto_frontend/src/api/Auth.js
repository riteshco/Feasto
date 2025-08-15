import { API_BASE_URL } from "./Config";

export async function loginUser(credentials) {
  const res = await fetch(`${API_BASE_URL}/auth`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },

    body: JSON.stringify(credentials),
    credentials: "include", // important if you want cookies (JWT) to be set
  });

  if (!res.ok) {
    throw new Error("Login failed");
  }

  return res.json();
}

export async function RegisterUser(credentials) {
  const res = await fetch(`${API_BASE_URL}/register`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },

    body: JSON.stringify(credentials),
  });

  if (!res.ok) {
    throw new Error("Register failed");
  }

  return res;
}

export const setCookie = (name, value , hours=1) => {
  const maxAge = hours * 60 * 60; 
  document.cookie = `${name}=${value || ""}; Max-Age=${maxAge}; path=/`;
};


