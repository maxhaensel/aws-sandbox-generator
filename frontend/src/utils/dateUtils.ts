function convertDate(str: string | undefined) {
  return new Date(str || '').toLocaleDateString('de-DE', {
    weekday: 'short',
    year: 'numeric',
    month: 'numeric',
    day: 'numeric',
  })
}

export { convertDate }
