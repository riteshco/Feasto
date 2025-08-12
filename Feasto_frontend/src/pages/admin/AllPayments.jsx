import { Navbar } from "@/components/Navbar"
import MainImage from "@/assets/food_home_image.jpg"
import { Input } from "@/components/ui/input"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle, } from "@/components/ui/card"
import { Label } from "@/components/ui/label"
import { Button } from "@/components/ui/button"
import { GetProducts } from "@/api/fetchAPI"
import { useEffect, useState } from "react"
import { AddToCartAPICall } from "@/api/Cart"
import { useNavigate } from "react-router-dom"
import { GetAllOrdersAPICall } from "@/api/fetchAPI"
import { AcceptOrderAPICall } from "@/api/AdminAction"
import { GetAllPaymentsAPICall } from "@/api/fetchAPI"

export function AllPaymentsPage() {

    const [AllPayments, setAllPayments] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        async function fetchAllOrders() {
            try {
                const allpymts = await GetAllPaymentsAPICall();
                setAllPayments(allpymts);
            } catch (err) {
                setError(err.message || "Failed to load products");
            } finally {
                setLoading(false);
            }
        }
        fetchAllOrders();

    }, []);


    if (loading) return <div>Loading orders...</div>;
    if (error) return <div>Error: {error}</div>;


    return (
        <>
            <Navbar user="admin" />
            <div className="relative w-full h-96 mt-16">
                <div className="Main_image h-full flex justify-center">
                    <img
                        src={MainImage}
                        alt="Main home image"
                        className="w-3/4 h-full rounded-3xl object-cover"
                    />
                </div>

                <div className="absolute inset-0 flex flex-col items-center justify-center bg-black/30 text-white p-4">
                    <h1 className="text-5xl font-bold mb-4">All Orders From all Users</h1>
                </div>
            </div>
            <div className="flex flex-col items-center cards w-full mt-8 gap-8">
                {AllPayments ? AllPayments.map((payment , index) => (
                    <Card className="w-3/4 flex">
                        <div className="cardinfo w-1/2">
                            <CardHeader>
                                <CardTitle className="text-3xl">Payment #{index+1}</CardTitle>
                                <CardDescription>
                                    Order ID : {payment.order_id}
                                </CardDescription>
                            </CardHeader>
                            <CardContent className="flex flex-col gap-4">
                                <div>
                                    CustomerID : {payment.user_id}
                                </div>
                                <div>
                                    Payment Status : {payment.payment_status === "completed" ? <span className="text-green-500">{payment.payment_status}</span> : <span className="text-red-500">{payment.payment_status}</span>}
                                </div>
                                <div>
                                    <span className="font-extrabold">Total Amount </span>: ${payment.Total_amount}
                                </div>
                            </CardContent>
                        </div>
                    </Card>
                ))
            : null}
            </div>
        </>
    )
}