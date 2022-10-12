import { withChildren, withClassName } from "../config/withs";

export default function Card ({ className = "", children }: CardProps) {
    return (
        <div className={"rounded-md shadow-md border border-gray-200 " + className}>
            {children}
        </div>
    )
}

export interface CardProps extends withClassName, withChildren {
}