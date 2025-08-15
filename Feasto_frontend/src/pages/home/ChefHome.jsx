import { Navbar } from "@/components/Navbar"
import MainImage from "@/assets/food_home_image.jpg"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle, } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { useEffect, useState } from "react"
import { GetAllOrdersAPICall } from "@/api/FetchAPI"
import { DeliverOrderAPICall } from "@/api/ChefAction"
import { toast , Toaster } from "sonner"

export function ChefHome() {

    const [AllOrders, setAllOrders] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        async function fetchAllOrders() {
            try {
                const allords = await GetAllOrdersAPICall();
                setAllOrders(allords);
            } catch (err) {
                setError(err.message || "Failed to load products");
            } finally {
                setLoading(false);
            }
        }
        fetchAllOrders();

    }, []);

    async function AskToDeliverOrder(OrderID){
        const message = await DeliverOrderAPICall(OrderID)
        toast(message, {
                action: {
                    label: "Ok",
                },
        })
        const allords = await GetAllOrdersAPICall();
        setAllOrders(allords);
        setLoading(false);
    }


    if (loading) return <div>Loading orders...</div>;
    if (error) return <div>Error: {error}</div>;


    return (
        <>
            <Navbar page="ChefHome" user="chef" />
            <div className="relative w-full h-96 mt-16">
                <Toaster position="top-center"/>
                <div className="Main_image h-full flex justify-center">
                    <img
                        src={MainImage}
                        alt="Main home image"
                        className="w-3/4 h-full rounded-3xl object-cover"
                    />
                </div>

                <div className="absolute inset-0 flex flex-col items-center justify-center bg-black/30 text-white p-4">
                    <h1 className="text-5xl font-bold mb-4">Welcome Back To Feasto</h1>
                </div>
            </div>
            <div className="flex flex-col items-center cards w-full mt-8 gap-8">
                {AllOrders ? AllOrders.map((order , index) => (
                    order.current_status !== "delivered" ?
                    <Card className="w-3/4 flex">
                        <div className="cardinfo w-1/2">
                            <CardHeader>
                                <CardTitle className="text-3xl">Order #{index+1}</CardTitle>
                                <CardDescription>
                                    Created at : {order.created_at.slice(0,10)}
                                </CardDescription>
                            </CardHeader>
                            <CardContent className="flex flex-col gap-4">
                                <div>
                                    Current Status : {order.current_status === "accepted" || order.current_status === "delivered" ? <span className="text-green-500">{order.current_status}</span> : <span className="text-red-500">{order.current_status}</span>}
                                </div>
                                <div>
                                    <span className="font-extrabold">Table Number </span>: {order.table_number}
                                </div>
                                <div>
                                    <span className="font-extrabold">Instructions</span> : {order.instructions}
                                </div>
                            </CardContent>
                        </div>
                        <div className="w-1/2">
                            <CardContent className="flex flex-col gap-4 items-center justify-center my-12 py-0">
                                {order.current_status === "accepted" ?
                                <div className="w-3/4 flex justify-center">
                                    <Button onClick={()=>{AskToDeliverOrder(order.id)}} className="w-1/2">Mark as delivered!</Button>
                                </div>
                                    :
                                <div className="w-3/4 flex justify-center" >
                                    <Button variant="outline" >Order not accepted/verified by Admin yet!</Button>
                                </div>}
                            </CardContent>
                        </div>
                    </Card>
                    : null 
                ))
            : null}
            </div>
        </>
    )
}