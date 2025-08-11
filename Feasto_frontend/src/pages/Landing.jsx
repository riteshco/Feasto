import { Navbar } from "@/components/Navbar"
import { AppWindowIcon, CodeIcon } from "lucide-react"

import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import {
    Tabs,
    TabsContent,
    TabsList,
    TabsTrigger,
} from "@/components/ui/tabs"

export function Landing() {
    return (
        <>
            <Navbar page="Landing" />
            <div className="flex items-center justify-center h-full w-full mt-16 flex-col gap-6 p-20">
                <Tabs defaultValue="Signup/Login" className="w-3/4">
                    <TabsList>
                        <TabsTrigger value="signup">Signup</TabsTrigger>
                        <TabsTrigger value="login">Login</TabsTrigger>
                    </TabsList>
                    <TabsContent value="signup">
                        <Card>
                            <CardHeader>
                                <CardTitle>Signup</CardTitle>
                                <CardDescription>
                                    Register yourself here on this Signup Page!
                                </CardDescription>
                            </CardHeader>
                            <CardContent className="grid gap-4">
                                <div className="grid gap-3">
                                    <Label htmlFor="tabs-demo-name">Username:</Label>
                                    <Input id="tabs-demo-name" placeholder="Your username goes here..." />
                                </div>
                                <div className="grid gap-3">
                                    <Label htmlFor="tabs-demo-username">Mobile Number:</Label>
                                    <Input id="tabs-demo-username" placeholder="10 digit number..." />
                                </div>
                                <div className="grid gap-3">
                                    <Label htmlFor="tabs-demo-username">Email</Label>
                                    <Input id="tabs-demo-username" placeholder="example@example.com" />
                                </div>
                                <div className="grid gap-3">
                                    <Label htmlFor="tabs-demo-username">Password</Label>
                                    <Input id="tabs-demo-username" placeholder="***********...." />
                                </div>
                            </CardContent>
                            <CardFooter>
                                <Button>Register</Button>
                            </CardFooter>
                        </Card>
                    </TabsContent>
                    <TabsContent value="login">
                        <Card>
                            <CardHeader>
                                <CardTitle>Login</CardTitle>
                                <CardDescription>
                                    Login to your account on Feasto.
                                </CardDescription>
                            </CardHeader>
                            <CardContent className="grid gap-6">
                                <div className="grid gap-3">
                                    <Label htmlFor="tabs-demo-name">Username:</Label>
                                    <Input id="tabs-demo-name" placeholder="Your username goes here..." />
                                </div>
                                <div className="grid gap-3">
                                    <Label htmlFor="tabs-demo-username">Email</Label>
                                    <Input id="tabs-demo-username" placeholder="example@example.com" />
                                </div>
                                <div className="grid gap-3">
                                    <Label htmlFor="tabs-demo-username">Password</Label>
                                    <Input id="tabs-demo-username" placeholder="***********...." />
                                </div>
                            </CardContent>
                            <CardFooter>
                                <Button>Login</Button>
                            </CardFooter>
                        </Card>
                    </TabsContent>
                </Tabs>
            </div>
        </>
    )
}