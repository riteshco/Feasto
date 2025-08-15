import { toast } from "sonner";
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

  const data = await res.json()

  if (!res.ok) {
    toast.error(data.message || "Something went wrong")
    return
  }
  return data;
}

export async function RegisterUser(credentials) {
  try{

    const res = await fetch(`${API_BASE_URL}/register`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      
      body: JSON.stringify(credentials),
    });
    
    const data = await res.json()

    if (!res.ok) {
      toast.error(data.message || "Something went wrong")
      return
    }

    toast.success(data.message)
  } catch (error){
    toast.error(error.message || "Network error")
  }
  }

export const setCookie = (name, value , hours=1) => {
  const maxAge = hours * 60 * 60; 
  document.cookie = `${name}=${value || ""}; Max-Age=${maxAge}; path=/`;
};


