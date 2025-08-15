import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle, } from "@/components/ui/card"
import { Navbar } from "@/components/Navbar";
import { Button } from "@/components/ui/button";
import { GetOrderPayment } from "@/api/FetchAPI";
import { PaymentDoneAPICall } from "@/api/Payment";
import { toast , Toaster } from "sonner";

export function OrderPaymentPage() {
  const { id } = useParams();
  const isValidId = /^\d+$/.test(id);
  const navigate = useNavigate()
  if (!isValidId) {
    return navigate("/not-found")
  }

  const [orderPaymentDetail, setPaymentDetail] = useState([]);

  useEffect(() => {
  async function AskToFetch() {
      let data = await GetOrderPayment(id);
      setPaymentDetail(data);
  }
  AskToFetch();
  }, []);

  function BackToOrders(){
    navigate("/orders");
  }

  async function PaymentDone(PaymentId){
    await PaymentDoneAPICall(PaymentId)
 
    let data = await GetOrderPayment(id);
    setPaymentDetail(data);
  }


    return (
        <>
            <Navbar page="OrderDetailPage" />
            <div className="relative w-full h-96 mt-16">
            <Toaster position="top-center" />
                <div className="absolute inset-0 flex flex-col items-center justify-center bg-black/30 text-white p-4">
                    <h1 className="text-5xl font-bold mb-4">Order #{id}</h1>
                </div>
            </div>

            <div className="flex flex-col items-center cards w-full gap-8">
                    <Card className="w-3/4 flex">
                        <div className="cardinfo w-3/4">
                            <CardHeader>
                                <CardTitle className="text-3xl">Payment of UserID : {orderPaymentDetail.user_id}</CardTitle>
                                <div className="total_amount">
                                Total Amount : ${orderPaymentDetail.Total_payment}
                                </div>
                                <div className="payment_status">
                                Payment Status : {orderPaymentDetail.payment_status === "completed" ? <span className="text-green-500">{orderPaymentDetail.payment_status}</span> : <span className="text-red-500">{orderPaymentDetail.payment_status}</span>}
                                </div>
                            </CardHeader>
                            <CardContent>
                                { orderPaymentDetail.payment_status !== "completed" ?
                                    <Button onClick={()=>{PaymentDone(orderPaymentDetail.id)}}>Mark payment as complete!</Button>
                                    :
                                    <Button variant="ghost">Payment done!</Button>
                                }
                            </CardContent>
                        </div>
                    </Card>
            <div className="m-8 mb-12">
            <Button onClick={BackToOrders}>Back To Orders Page</Button>
            </div>
            </div>
        </>
    )
}