/*
 * Brute Force Detector
 *
 * You're analyzing authentication logs. Each log entry is a string like:
 *
 *   192.168.1.1 LOGIN_FAIL user=admin
 *   192.168.1.1 LOGIN_SUCCESS user=admin
 *   10.0.0.5 LOGIN_FAIL user=root
 *
 * An IP is suspicious if it has 3 or more consecutive failures without
 * a success in between. Return a list of suspicious IPs, sorted alphabetically.
 *
 * A success resets that IP's failure count to zero.
 * Empty list returns empty list.
 *
 * You do not need to validate the IP addresses.
 */

function detectBruteForce(logs) {
}

module.exports = { detectBruteForce };