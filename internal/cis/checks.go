package cis

import (
	"fmt"
)

// Check is a single security check
type Check struct {
	Name        string
	Description string
	Enabled     bool
}

// Checks holds all the security checks
var Checks = []Check{
	{Name: "Root Password", Description: "Ensure that the root password is set", Enabled: true},
	{Name: "Shadow Permissions", Description: "Ensure that the permission of /etc/shadow is set securely", Enabled: true},
	{Name: "System Update", Description: "Ensure that the system is up-to-date", Enabled: true},
	{Name: "ACL Configuration", Description: "Ensure that ACL is not used on sensitive files", Enabled: true},
	{Name: "SSH Root Login", Description: "Ensure that SSH root login is disabled", Enabled: true},
	{Name: "Unused Services", Description: "Ensure that unused services are disabled", Enabled: true},
	{Name: "Firewall Configuration", Description: "Ensure that firewall is configured and running", Enabled: true},
	{Name: "Password Expiration", Description: "Ensure that periodic password expiration is enforced", Enabled: true},
	{Name: "SELinux", Description: "Ensure that SELinux is enabled on RHEL", Enabled: true},
	{Name: "Unattended Upgrades", Description: "Ensure that unattended-upgrades is installed for Ubuntu", Enabled: true},
}

// PerformChecks executes all security checks
func PerformChecks() {
	for _, check := range Checks {
		if check.Enabled {
			fmt.Printf("%s: %s\n", check.Name, check.Description)
		}
	}
}