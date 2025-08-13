import { useEffect, useState } from "react";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle, } from "@/components/ui/card"
import { Navbar } from "@/components/Navbar";
import { Button } from "@/components/ui/button";
import { fetchUserOrder } from "@/api/fetchAPI";
import { useNavigate } from "react-router-dom";

export function UserOrderPage() {
  const [orders, setOrders] = useState([]);
  const navigate = useNavigate()

  useEffect(() => {
  async function AskToFetch() {
      let data = await fetchUserOrder();
      setOrders(data);
      console.log(data)
  }
  AskToFetch();
  }, []);

  async function GoToOrderDetail(OrderID){
    navigate(`/order/items/${OrderID}`)
  }

  async function GoToPayment(OrderID){
    navigate(`/order/payment/${OrderID}`)
  }


    return (
        <>
            <Navbar page="UserOrderPage" />
            <div className="relative w-full h-96 mt-16">
                <div className="absolute inset-0 flex flex-col items-center justify-center bg-black/30 text-white p-4">
                    <h1 className="text-5xl font-bold mb-4">Your Orders!</h1>
                </div>
            </div>

            <div className="flex flex-col items-center cards w-full mt-8 gap-8">
                {orders ? orders.map((order , index) => (

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
                                <div className="w-3/4 flex justify-center">
                                    <Button onClick={()=>{GoToOrderDetail(order.id)}} variant="outline" className="w-1/2">View Details</Button>
                                </div>
                                <div className="w-3/4 flex justify-center">
                                    <Button onClick={()=>{GoToPayment(order.id)}} className="w-1/2">Payment</Button>
                                </div>
                                {order.current_status === "accepted" || order.current_status === "delivered" ?
                                <div className="w-3/4 flex justify-center">
                                    <Button variant="outline" className="w-1/2">Bill</Button>
                                </div>
                                :null}
                            </CardContent>
                        </div>
                    </Card>
                ))
            : null}
            </div>
        </>
    )
}