resource "cloudflare_zone_lockdown" "%[1]s" {
    zone_id = "%[2]s"
    urls    = ["%[6]s"]
    configurations = [
      {
        target = "%[7]s"
        value  = "%[8]s"
      }
    ]
}
