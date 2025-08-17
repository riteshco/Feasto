import { useEffect, useState } from "react";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle, } from "@/components/ui/card"
import { Navbar } from "@/components/Navbar";
import { fetchChefDeliveredOrders } from "@/api/FetchAPI";


export function DeliveredOrdersPage() {
  const [pastOrders, setDeliveredOrders] = useState([]);

  useEffect(() => {
  async function AskToFetch() {
      let data = await fetchChefDeliveredOrders();
      setDeliveredOrders(data);
  }
  AskToFetch();
  }, []);

    return (
        <>
            <Navbar page="DeliveredOrdersPage" />
            <div className="relative w-full h-96 mt-16">
                <div className="absolute inset-0 flex flex-col items-center justify-center bg-black/30 text-white p-4">
                    <h1 className="text-5xl font-bold mb-4">Your Delivered Orders!</h1>
                </div>
            </div>

            <div className="flex flex-col items-center cards w-full mt-8 gap-8">
                {pastOrders ? pastOrders.map((order , index) => (

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
                    </Card>
                ))
            : null}
            </div>
        </>
    )
}