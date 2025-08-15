import { API_BASE_URL } from "./Config";

export const GetProducts = async () => {
    try {
        const response = await fetch(
            `${API_BASE_URL}/all-products`
        ,{
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        });
        const products = await response.json()
        return products
    }
    catch (error) {
        alert(error.message);
    }
}

export const fetchUserOrder = async () => {
    try {
        const response = await fetch(
            `${API_BASE_URL}/orders`
        ,{
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        });
        const orders = await response.json()
        return orders
    } catch (error){
        alert(error.message)
    }
}

export const GetAllOrdersAPICall = async () => {
    try {
        const response = await fetch(
            `${API_BASE_URL}/all-orders`
        ,{
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        });
        const orders = await response.json()
        return orders
    } catch (error){
        alert(error.message)
    }
}

export const GetAllPaymentsAPICall = async () => {
    try {
        const response = await fetch(
            `${API_BASE_URL}/all-payments`
        ,{
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        });
        const payments = await response.json()
        return payments
    } catch (error){
        alert(error.message)
    }
}

export const fetchUserPastOrder = async () => {
    try {
        const response = await fetch(
            `${API_BASE_URL}/past-orders`
        ,{
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        });
        const orders = await response.json()
        return orders
    } catch (error){
        alert(error.message)
    }
}

export const fetchUsers = async () => {
    try {
        const response = await fetch(
            `${API_BASE_URL}/users`
        ,{
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        });
        const users = await response.json()
        return users
    } catch (error){
        alert(error.message)
    }
}

export const GetUserByID = async (id) => {
    try {
        const response = await fetch(
            `${API_BASE_URL}/user/${id}`
        );
        user = response.json()
    }
    catch (error) {
        alert(error.message);
    }
}
export async function GetOrderPayment(OrderId) {
    try{
        const res = await fetch(`${API_BASE_URL}/order/payment/${OrderId}`, {
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
        });
        const data = await res.json();
        return data
    } catch (error){
        alert(error.message)
    }
}

export async function GetOrderDetail(OrderId) {
  const res = await fetch(`${API_BASE_URL}/order/items/${OrderId}`, {
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });
  const data = await res.json();
  if (data.orders) {

    return data.orders.map(order => {
      const product = data.products.find(p => p.id === order.product_id);
      return {
        orderId: order.id,
        product_name: product ? product.product_name : "Unknown Product",
        quantity: order.quantity,
        price: product.price
      };
    });
  } else {
    return null
  }
}