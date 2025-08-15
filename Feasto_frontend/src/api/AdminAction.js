import { toast } from "sonner";
import { API_BASE_URL } from "./Config";

export async function ChangeUserRoleAPICall(UserId , newRole) {
    try {
    const res = await fetch(`${API_BASE_URL}/edit-user-role/${UserId}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({"user_role":newRole}),
      credentials: "include",
    });
    if (!res.ok) {const data = await res.json();toast.error(data.message || "Something went wrong"); return}
    toast.success(`Successfully Changed role to ${newRole} of User #${UserId}`)
  } catch (err) {
    toast.error(err.message)
  }
}

export async function AcceptOrderAPICall(OrderId) {
    try {
    const res = await fetch(`${API_BASE_URL}/gen-bill/${OrderId}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });
    if (!res.ok) {const data = await res.json();toast.error(data.message || "Something went wrong"); return}
    toast.success(`Successfully accepted the Order#${OrderId}`)
  } catch (err) {
    toast.error(err.message)
  }
}