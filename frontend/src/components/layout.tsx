interface Props {
  children: React.ReactNode
}

function Layout({ children }: Props) {
  return (
    <>
      <div className="flex justify-center">
        <div className="m-16">{children}</div>
      </div>
    </>
  )
}

export { Layout }
