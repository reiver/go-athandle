package athandle

// Resolve resolved an AT-protocol handle to a DID.
//
// For example:
//
//	did, err := athandle.Resolve("reiver.bsky.social")
func Resolve(handle string) (did string, err error) {

	{
		results, err := TXTLookup(handle)
		if nil != err {
			switch err.(type) {
			case DIDTXTNotFoundError:
				// nothing here
			case *DIDTXTNotFoundError:
				// nothing here
			default:
				return "", err
			}
		}
		if 0 < len(results) {
			return results[0], nil
		}
	}

	{
		result, err := WellKnownLookup(handle)
		if nil != err {
			return "", err
		}

		return result, nil
	}
}
