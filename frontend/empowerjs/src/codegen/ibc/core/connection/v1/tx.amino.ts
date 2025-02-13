import { AminoMsg } from "@cosmjs/amino";
import { AminoHeight, omitDefault } from "../../../../helpers";
import { MsgConnectionOpenInit, MsgConnectionOpenTry, MsgConnectionOpenAck, MsgConnectionOpenConfirm } from "./tx";
export interface MsgConnectionOpenInitAminoType extends AminoMsg {
  type: "cosmos-sdk/MsgConnectionOpenInit";
  value: {
    client_id: string;
    counterparty: {
      client_id: string;
      connection_id: string;
      prefix: {
        key_prefix: Uint8Array;
      };
    };
    version: {
      identifier: string;
      features: string[];
    };
    delay_period: string;
    signer: string;
  };
}
export interface MsgConnectionOpenTryAminoType extends AminoMsg {
  type: "cosmos-sdk/MsgConnectionOpenTry";
  value: {
    client_id: string;
    previous_connection_id: string;
    client_state: {
      type_url: string;
      value: Uint8Array;
    };
    counterparty: {
      client_id: string;
      connection_id: string;
      prefix: {
        key_prefix: Uint8Array;
      };
    };
    delay_period: string;
    counterparty_versions: {
      identifier: string;
      features: string[];
    }[];
    proof_height: AminoHeight;
    proof_init: Uint8Array;
    proof_client: Uint8Array;
    proof_consensus: Uint8Array;
    consensus_height: AminoHeight;
    signer: string;
    host_consensus_state_proof: Uint8Array;
  };
}
export interface MsgConnectionOpenAckAminoType extends AminoMsg {
  type: "cosmos-sdk/MsgConnectionOpenAck";
  value: {
    connection_id: string;
    counterparty_connection_id: string;
    version: {
      identifier: string;
      features: string[];
    };
    client_state: {
      type_url: string;
      value: Uint8Array;
    };
    proof_height: AminoHeight;
    proof_try: Uint8Array;
    proof_client: Uint8Array;
    proof_consensus: Uint8Array;
    consensus_height: AminoHeight;
    signer: string;
    host_consensus_state_proof: Uint8Array;
  };
}
export interface MsgConnectionOpenConfirmAminoType extends AminoMsg {
  type: "cosmos-sdk/MsgConnectionOpenConfirm";
  value: {
    connection_id: string;
    proof_ack: Uint8Array;
    proof_height: AminoHeight;
    signer: string;
  };
}
export const AminoConverter = {
  "/ibc.core.connection.v1.MsgConnectionOpenInit": {
    aminoType: "cosmos-sdk/MsgConnectionOpenInit",
    toAmino: ({
      clientId,
      counterparty,
      version,
      delayPeriod,
      signer
    }: MsgConnectionOpenInit): MsgConnectionOpenInitAminoType["value"] => {
      return {
        client_id: clientId,
        counterparty: {
          client_id: counterparty.clientId,
          connection_id: counterparty.connectionId,
          prefix: {
            key_prefix: counterparty.prefix.keyPrefix
          }
        },
        version: {
          identifier: version.identifier,
          features: version.features
        },
        delay_period: delayPeriod.toString(),
        signer
      };
    },
    fromAmino: ({
      client_id,
      counterparty,
      version,
      delay_period,
      signer
    }: MsgConnectionOpenInitAminoType["value"]): MsgConnectionOpenInit => {
      return {
        clientId: client_id,
        counterparty: {
          clientId: counterparty.client_id,
          connectionId: counterparty.connection_id,
          prefix: {
            keyPrefix: counterparty.prefix.key_prefix
          }
        },
        version: {
          identifier: version.identifier,
          features: version.features
        },
        delayPeriod: BigInt(delay_period),
        signer
      };
    }
  },
  "/ibc.core.connection.v1.MsgConnectionOpenTry": {
    aminoType: "cosmos-sdk/MsgConnectionOpenTry",
    toAmino: ({
      clientId,
      previousConnectionId,
      clientState,
      counterparty,
      delayPeriod,
      counterpartyVersions,
      proofHeight,
      proofInit,
      proofClient,
      proofConsensus,
      consensusHeight,
      signer,
      hostConsensusStateProof
    }: MsgConnectionOpenTry): MsgConnectionOpenTryAminoType["value"] => {
      return {
        client_id: clientId,
        previous_connection_id: previousConnectionId,
        client_state: {
          type_url: clientState.typeUrl,
          value: clientState.value
        },
        counterparty: {
          client_id: counterparty.clientId,
          connection_id: counterparty.connectionId,
          prefix: {
            key_prefix: counterparty.prefix.keyPrefix
          }
        },
        delay_period: delayPeriod.toString(),
        counterparty_versions: counterpartyVersions.map(el0 => ({
          identifier: el0.identifier,
          features: el0.features
        })),
        proof_height: proofHeight ? {
          revision_height: omitDefault(proofHeight.revisionHeight)?.toString(),
          revision_number: omitDefault(proofHeight.revisionNumber)?.toString()
        } : {},
        proof_init: proofInit,
        proof_client: proofClient,
        proof_consensus: proofConsensus,
        consensus_height: consensusHeight ? {
          revision_height: omitDefault(consensusHeight.revisionHeight)?.toString(),
          revision_number: omitDefault(consensusHeight.revisionNumber)?.toString()
        } : {},
        signer,
        host_consensus_state_proof: hostConsensusStateProof
      };
    },
    fromAmino: ({
      client_id,
      previous_connection_id,
      client_state,
      counterparty,
      delay_period,
      counterparty_versions,
      proof_height,
      proof_init,
      proof_client,
      proof_consensus,
      consensus_height,
      signer,
      host_consensus_state_proof
    }: MsgConnectionOpenTryAminoType["value"]): MsgConnectionOpenTry => {
      return {
        clientId: client_id,
        previousConnectionId: previous_connection_id,
        clientState: {
          typeUrl: client_state.type_url,
          value: client_state.value
        },
        counterparty: {
          clientId: counterparty.client_id,
          connectionId: counterparty.connection_id,
          prefix: {
            keyPrefix: counterparty.prefix.key_prefix
          }
        },
        delayPeriod: BigInt(delay_period),
        counterpartyVersions: counterparty_versions.map(el0 => ({
          identifier: el0.identifier,
          features: el0.features
        })),
        proofHeight: proof_height ? {
          revisionHeight: BigInt(proof_height.revision_height || "0", true),
          revisionNumber: BigInt(proof_height.revision_number || "0", true)
        } : undefined,
        proofInit: proof_init,
        proofClient: proof_client,
        proofConsensus: proof_consensus,
        consensusHeight: consensus_height ? {
          revisionHeight: BigInt(consensus_height.revision_height || "0", true),
          revisionNumber: BigInt(consensus_height.revision_number || "0", true)
        } : undefined,
        signer,
        hostConsensusStateProof: host_consensus_state_proof
      };
    }
  },
  "/ibc.core.connection.v1.MsgConnectionOpenAck": {
    aminoType: "cosmos-sdk/MsgConnectionOpenAck",
    toAmino: ({
      connectionId,
      counterpartyConnectionId,
      version,
      clientState,
      proofHeight,
      proofTry,
      proofClient,
      proofConsensus,
      consensusHeight,
      signer,
      hostConsensusStateProof
    }: MsgConnectionOpenAck): MsgConnectionOpenAckAminoType["value"] => {
      return {
        connection_id: connectionId,
        counterparty_connection_id: counterpartyConnectionId,
        version: {
          identifier: version.identifier,
          features: version.features
        },
        client_state: {
          type_url: clientState.typeUrl,
          value: clientState.value
        },
        proof_height: proofHeight ? {
          revision_height: omitDefault(proofHeight.revisionHeight)?.toString(),
          revision_number: omitDefault(proofHeight.revisionNumber)?.toString()
        } : {},
        proof_try: proofTry,
        proof_client: proofClient,
        proof_consensus: proofConsensus,
        consensus_height: consensusHeight ? {
          revision_height: omitDefault(consensusHeight.revisionHeight)?.toString(),
          revision_number: omitDefault(consensusHeight.revisionNumber)?.toString()
        } : {},
        signer,
        host_consensus_state_proof: hostConsensusStateProof
      };
    },
    fromAmino: ({
      connection_id,
      counterparty_connection_id,
      version,
      client_state,
      proof_height,
      proof_try,
      proof_client,
      proof_consensus,
      consensus_height,
      signer,
      host_consensus_state_proof
    }: MsgConnectionOpenAckAminoType["value"]): MsgConnectionOpenAck => {
      return {
        connectionId: connection_id,
        counterpartyConnectionId: counterparty_connection_id,
        version: {
          identifier: version.identifier,
          features: version.features
        },
        clientState: {
          typeUrl: client_state.type_url,
          value: client_state.value
        },
        proofHeight: proof_height ? {
          revisionHeight: BigInt(proof_height.revision_height || "0", true),
          revisionNumber: BigInt(proof_height.revision_number || "0", true)
        } : undefined,
        proofTry: proof_try,
        proofClient: proof_client,
        proofConsensus: proof_consensus,
        consensusHeight: consensus_height ? {
          revisionHeight: BigInt(consensus_height.revision_height || "0", true),
          revisionNumber: BigInt(consensus_height.revision_number || "0", true)
        } : undefined,
        signer,
        hostConsensusStateProof: host_consensus_state_proof
      };
    }
  },
  "/ibc.core.connection.v1.MsgConnectionOpenConfirm": {
    aminoType: "cosmos-sdk/MsgConnectionOpenConfirm",
    toAmino: ({
      connectionId,
      proofAck,
      proofHeight,
      signer
    }: MsgConnectionOpenConfirm): MsgConnectionOpenConfirmAminoType["value"] => {
      return {
        connection_id: connectionId,
        proof_ack: proofAck,
        proof_height: proofHeight ? {
          revision_height: omitDefault(proofHeight.revisionHeight)?.toString(),
          revision_number: omitDefault(proofHeight.revisionNumber)?.toString()
        } : {},
        signer
      };
    },
    fromAmino: ({
      connection_id,
      proof_ack,
      proof_height,
      signer
    }: MsgConnectionOpenConfirmAminoType["value"]): MsgConnectionOpenConfirm => {
      return {
        connectionId: connection_id,
        proofAck: proof_ack,
        proofHeight: proof_height ? {
          revisionHeight: BigInt(proof_height.revision_height || "0", true),
          revisionNumber: BigInt(proof_height.revision_number || "0", true)
        } : undefined,
        signer
      };
    }
  }
};