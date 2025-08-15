import { API_BASE_URL } from "./Config";

export async function AddChangeRequest(role) {
  try {
    const res = await fetch(`${API_BASE_URL}/change_role_request/${role}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });
    if (!res.ok) throw new Error("Failed to request role change!");
    let message = `Request to change role sent`
    return message

  } catch (err) {
    return err.message
  }
}