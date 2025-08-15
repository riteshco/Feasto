import { toast } from "sonner";
import { API_BASE_URL } from "./Config";

export async function AddChangeRequest(role) {
  try {
    const res = await fetch(`${API_BASE_URL}/change_role_request/${role}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });
    if (!res.ok) {const data = await res.json();toast.error(data.message || "Failed to request role change!"); return}
    toast.success(`Request to change role sent`)

  } catch (err) {
    toast.error(err.message)
  }
}