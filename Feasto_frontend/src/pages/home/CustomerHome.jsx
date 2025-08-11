import { Navbar } from "@/components/Navbar"
import MainImage from "@/assets/food_home_image.jpg"
import { Input } from "@/components/ui/input"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle, } from "@/components/ui/card"
import { Label } from "@/components/ui/label"
import { Button } from "@/components/ui/button"

export function CustomerHome() {
    return (
        <>
            <Navbar page="CustomerHome" />
            <div className="relative w-full h-96 mt-16">
                {/* Background Image */}
                <div className="Main_image h-full flex justify-center">
                    <img
                        src={MainImage}
                        alt="Main home image"
                        className="w-3/4 h-full rounded-3xl object-cover"
                    />
                </div>

                {/* Overlay Content */}
                <div className="absolute inset-0 flex flex-col items-center justify-center bg-black/30 text-white p-4">
                    <h1 className="text-5xl font-bold mb-4">Welcome to Feasto</h1>
                    <form action="" className="flex justify-center gap-4 w-full">
                    <Input className="w-1/2 font-extrabold text-2xl bg-white/15" placeholder="Search products..." />
                    <Button>Search</Button>
                    </form>
                </div>
            </div>

            <div className="flex flex-col items-center cards w-full mt-8">

                <Card className="w-3/4 flex">
                    <div className="cardimage w-1/4 flex justify-center items-center px-4">
                        <img src={MainImage} alt="Product_image" className="w-full rounded-2xl"/>
                    </div>
                    <div className="cardinfo w-3/4">

                        <CardHeader>
                            <CardTitle>Login to your account</CardTitle>
                            <CardDescription>
                                Enter your email below to login to your account
                            </CardDescription>
                            <Button variant="link">Sign Up</Button>
                        </CardHeader>
                        <CardContent>
                            <form>
                                <div className="flex flex-col gap-6">
                                    <div className="grid gap-2">
                                        <Label htmlFor="email">Email</Label>
                                        <Input
                                            id="email"
                                            type="email"
                                            placeholder="m@example.com"
                                            required
                                        />
                                    </div>
                                    <div className="grid gap-2">
                                        <div className="flex items-center">
                                            <Label htmlFor="password">Password</Label>
                                            <a
                                                href="#"
                                                className="ml-auto inline-block text-sm underline-offset-4 hover:underline"
                                            >
                                                Forgot your password?
                                            </a>
                                        </div>
                                        <Input id="password" type="password" required />
                                    </div>
                                </div>
                            </form>
                        </CardContent>
                    </div>
                </Card>
                
            </div>
        </>
    )
}