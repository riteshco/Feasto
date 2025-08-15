import { API_BASE_URL } from "./Config";

export async function ChangeUserRoleAPICall(UserId , newRole) {
    try {
    const res = await fetch(`${API_BASE_URL}/edit-user-role/${UserId}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({"user_role":newRole}),
      credentials: "include",
    });
    if (!res.ok) throw new Error("Failed to change the role");
    alert(`Changed user #${UserId} to chef!`);
  } catch (err) {
    alert(err.message);
  }
}

export async function AcceptOrderAPICall(OrderId) {
    try {
    const res = await fetch(`${API_BASE_URL}/gen-bill/${OrderId}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });
    if (!res.ok) throw new Error("Failed to accept the order");
    alert(`Accepted Order #${OrderId} from the customer!`);
  } catch (err) {
    alert(err.message);
  }
}