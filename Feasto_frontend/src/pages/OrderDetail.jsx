import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle, } from "@/components/ui/card"
import { Navbar } from "@/components/Navbar";
import { Button } from "@/components/ui/button";
import { GetOrderDetail } from "@/api/FetchAPI";

export function OrderDetailPage() {
  const { id } = useParams();
  const isValidId = /^\d+$/.test(id);
  const navigate = useNavigate()
  if (!isValidId) {
    return navigate("/not-found")
  }

  const [orderDetail, setDetail] = useState([]);

  useEffect(() => {
  async function AskToFetch() {
      let merged = await GetOrderDetail(id);
      setDetail(merged);
  }
  AskToFetch();
  }, []);

  function BackToOrders(){
    navigate("/orders");
  }


    return (
        <>
            <Navbar page="OrderDetailPage" />
            <div className="relative w-full h-96 mt-16">
                <div className="absolute inset-0 flex flex-col items-center justify-center bg-black/30 text-white p-4">
                    <h1 className="text-5xl font-bold mb-4">Order #{id}</h1>
                </div>
            </div>

            <div className="flex flex-col items-center cards w-full gap-8">
                {orderDetail ? orderDetail.map((orderItem) => (

                    <Card className="w-3/4 flex">
                        <div className="cardinfo w-3/4">
                            <CardHeader>
                                <CardTitle className="text-3xl">{orderItem.product_name}</CardTitle>
                                <div className="quantity">
                                Quantity : {orderItem.quantity}
                                </div>
                                <div className="price">
                                Price : ${orderItem.price}
                                </div>
                            </CardHeader>
                            <CardContent>
                            </CardContent>
                        </div>
                    </Card>
                ))
            : null}
            <div className="m-8 mb-12">
            <Button onClick={BackToOrders}>Back To Orders Page</Button>
            </div>
            </div>
        </>
    )
}