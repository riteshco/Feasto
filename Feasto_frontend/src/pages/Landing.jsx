import { Navbar } from "@/components/Navbar"
import { useState } from "react"
import { useNavigate } from "react-router-dom"
import { loginUser , setCookie , RegisterUser} from "@/api/Auth"
import { Button } from "@/components/ui/button"
import { Navigate } from "react-router-dom"
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
import { getUserFromToken } from "@/utils/Auth"
import { toast , Toaster } from "sonner"

export function Landing() {
    const user = getUserFromToken()
    
    if(user) {
        if(user.user_role === "admin"){
            return <Navigate to="/users" replace />;
        }
        return <Navigate to="/home" replace />;
    }

    const [mobile_number, setMobileNumber] = useState("");
    async function handleRegisterSubmit(e) {
        e.preventDefault();
        try {
        await RegisterUser({ username , mobile_number , email , password });
        navigate("/home")
        } catch (err) {
            console.error(err);
        }
    }

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [email   , setEmail]    = useState("");
    const navigate = useNavigate();

    async function handleSubmit(e) {
        e.preventDefault();
        try {
        const data = await loginUser({ username, email , password });
        setCookie("auth_token",data.token);
        if(data.role === "admin"){
            navigate("/users")
        } else {
            navigate("/home");
        }
        } catch (err) {
            console.error(err);
        }
    }

    return (
        <>
            <Navbar page="Landing" />
            <div className="flex items-center justify-center h-full w-full mt-16 flex-col gap-6 p-20">
                <Toaster position="top-center"/>
                <Tabs defaultValue="login" className="w-3/4">
                <div className="flex justify-center">
                    <TabsList className="">
                        <TabsTrigger value="signup">Signup</TabsTrigger>
                        <TabsTrigger value="login">Login</TabsTrigger>
                    </TabsList>
                </div>
                    <TabsContent value="signup">
                        <Card>
                            <CardHeader>
                                <CardTitle>Signup</CardTitle>
                                <CardDescription>
                                    Register yourself here on this Signup Page!
                                </CardDescription>
                            </CardHeader>
                            <form onSubmit={handleRegisterSubmit}>
                            <CardContent className="grid gap-4">
                                <div className="grid gap-3">
                                    <Label htmlFor="tabs-demo-name">Username:</Label>
                                    <Input value={username} type="text" maxLength={15} onChange={(e) => setUsername(e.target.value)} id="tabs-demo-name" placeholder="Your username goes here..." required />
                                </div>
                                <div className="grid gap-3">
                                    <Label htmlFor="tabs-demo-username">Mobile Number:</Label>
                                    <Input value={mobile_number} type="text" maxLength={10} minLength={10} onChange={(e) => setMobileNumber(e.target.value.replace(/\D/g, ""))} id="tabs-demo-username" placeholder="10 digit number..." required />
                                </div>
                                <div className="grid gap-3">
                                    <Label htmlFor="tabs-demo-username">Email</Label>
                                    <Input value={email} type="email" onChange={(e) => setEmail(e.target.value)} id="tabs-demo-username" placeholder="example@example.com" required/>
                                </div>
                                <div className="grid gap-3">
                                    <Label htmlFor="tabs-demo-username">Password</Label>
                                    <Input value={password} type="password" onChange={(e) => setPassword(e.target.value)} id="tabs-demo-username" placeholder="***********...." required/>
                                </div>
                            </CardContent>
                            <CardFooter>
                                <Button>Register</Button>
                            </CardFooter>
                            </form>
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
                                <form onSubmit={handleSubmit}>
                            <CardContent className="grid gap-6">
                                <div className="grid gap-3">
                                    <Label htmlFor="tabs-demo-name">Username:</Label>
                                    <Input value={username} onChange={(e) => setUsername(e.target.value)} id="tabs-demo-name" placeholder="Your username goes here..." required />
                                </div>
                                <div className="grid gap-3">
                                    <Label htmlFor="tabs-demo-username">Email</Label>
                                    <Input value={email} type="email" onChange={(e) => setEmail(e.target.value)} id="tabs-demo-username" placeholder="example@example.com" required />
                                </div>
                                <div className="grid gap-3">
                                    <Label htmlFor="tabs-demo-username">Password</Label>
                                    <Input value={password} type="password" onChange={(e) => setPassword(e.target.value)} id="tabs-demo-username" placeholder="***********...." required />
                                </div>
                            </CardContent>
                            <CardFooter>
                                <Button type="submit">Login</Button>
                            </CardFooter>
                                </form>
                        </Card>
                    </TabsContent>
                </Tabs>
            </div>
        </>
    )
}