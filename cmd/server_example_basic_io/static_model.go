package main

import (
	"fmt"

	libiec61850go "github.com/iotxfoundry/libiec61850-go"
)

var (
	iedModelds_GenericIO_LLN0_Events                    = libiec61850go.NewDataSet()
	iedModelds_GenericIO_LLN0_Events2                   = libiec61850go.NewDataSet()
	iedModelds_GenericIO_LLN0_Measurements              = libiec61850go.NewDataSet()
	iedModelds_GenericIO_LLN0_Events_fcda0              = libiec61850go.NewDataSetEntry()
	iedModelds_GenericIO_LLN0_Events_fcda1              = libiec61850go.NewDataSetEntry()
	iedModelds_GenericIO_LLN0_Events_fcda2              = libiec61850go.NewDataSetEntry()
	iedModelds_GenericIO_LLN0_Events_fcda3              = libiec61850go.NewDataSetEntry()
	iedModelds_GenericIO_LLN0_Events2_fcda0             = libiec61850go.NewDataSetEntry()
	iedModelds_GenericIO_LLN0_Events2_fcda1             = libiec61850go.NewDataSetEntry()
	iedModelds_GenericIO_LLN0_Events2_fcda2             = libiec61850go.NewDataSetEntry()
	iedModelds_GenericIO_LLN0_Events2_fcda3             = libiec61850go.NewDataSetEntry()
	iedModelds_GenericIO_LLN0_Measurements_fcda0        = libiec61850go.NewDataSetEntry()
	iedModelds_GenericIO_LLN0_Measurements_fcda1        = libiec61850go.NewDataSetEntry()
	iedModelds_GenericIO_LLN0_Measurements_fcda2        = libiec61850go.NewDataSetEntry()
	iedModelds_GenericIO_LLN0_Measurements_fcda3        = libiec61850go.NewDataSetEntry()
	iedModelds_GenericIO_LLN0_Measurements_fcda4        = libiec61850go.NewDataSetEntry()
	iedModelds_GenericIO_LLN0_Measurements_fcda5        = libiec61850go.NewDataSetEntry()
	iedModelds_GenericIO_LLN0_Measurements_fcda6        = libiec61850go.NewDataSetEntry()
	iedModelds_GenericIO_LLN0_Measurements_fcda7        = libiec61850go.NewDataSetEntry()
	iedModel                                            = libiec61850go.NewIedModel()
	iedModel_GenericIO                                  = libiec61850go.NewLogicalDevice()
	iedModel_GenericIO_LLN0                             = libiec61850go.NewLogicalNode()
	iedModel_GenericIO_LLN0_Mod                         = libiec61850go.NewDataObject()
	iedModel_GenericIO_LLN0_Mod_stVal                   = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LLN0_Mod_q                       = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LLN0_Mod_t                       = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LLN0_Mod_ctlModel                = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LLN0_Beh                         = libiec61850go.NewDataObject()
	iedModel_GenericIO_LLN0_Beh_stVal                   = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LLN0_Beh_q                       = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LLN0_Beh_t                       = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LLN0_Health                      = libiec61850go.NewDataObject()
	iedModel_GenericIO_LLN0_Health_stVal                = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LLN0_Health_q                    = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LLN0_Health_t                    = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LLN0_NamPlt                      = libiec61850go.NewDataObject()
	iedModel_GenericIO_LLN0_NamPlt_vendor               = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LLN0_NamPlt_swRev                = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LLN0_NamPlt_d                    = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LLN0_NamPlt_configRev            = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LLN0_NamPlt_ldNs                 = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LPHD1                            = libiec61850go.NewLogicalNode()
	iedModel_GenericIO_LPHD1_PhyNam                     = libiec61850go.NewDataObject()
	iedModel_GenericIO_LPHD1_PhyNam_vendor              = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LPHD1_PhyHealth                  = libiec61850go.NewDataObject()
	iedModel_GenericIO_LPHD1_PhyHealth_stVal            = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LPHD1_PhyHealth_q                = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LPHD1_PhyHealth_t                = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LPHD1_Proxy                      = libiec61850go.NewDataObject()
	iedModel_GenericIO_LPHD1_Proxy_stVal                = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LPHD1_Proxy_q                    = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_LPHD1_Proxy_t                    = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1                            = libiec61850go.NewLogicalNode()
	iedModel_GenericIO_GGIO1_Mod                        = libiec61850go.NewDataObject()
	iedModel_GenericIO_GGIO1_Mod_stVal                  = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Mod_q                      = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Mod_t                      = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Mod_ctlModel               = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Beh                        = libiec61850go.NewDataObject()
	iedModel_GenericIO_GGIO1_Beh_stVal                  = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Beh_q                      = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Beh_t                      = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Health                     = libiec61850go.NewDataObject()
	iedModel_GenericIO_GGIO1_Health_stVal               = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Health_q                   = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Health_t                   = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_NamPlt                     = libiec61850go.NewDataObject()
	iedModel_GenericIO_GGIO1_NamPlt_vendor              = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_NamPlt_swRev               = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_NamPlt_d                   = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_AnIn1                      = libiec61850go.NewDataObject()
	iedModel_GenericIO_GGIO1_AnIn1_mag                  = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_AnIn1_mag_f                = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_AnIn1_q                    = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_AnIn1_t                    = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_AnIn2                      = libiec61850go.NewDataObject()
	iedModel_GenericIO_GGIO1_AnIn2_mag                  = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_AnIn2_mag_f                = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_AnIn2_q                    = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_AnIn2_t                    = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_AnIn3                      = libiec61850go.NewDataObject()
	iedModel_GenericIO_GGIO1_AnIn3_mag                  = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_AnIn3_mag_f                = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_AnIn3_q                    = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_AnIn3_t                    = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_AnIn4                      = libiec61850go.NewDataObject()
	iedModel_GenericIO_GGIO1_AnIn4_mag                  = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_AnIn4_mag_f                = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_AnIn4_q                    = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_AnIn4_t                    = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO1                     = libiec61850go.NewDataObject()
	iedModel_GenericIO_GGIO1_SPCSO1_origin              = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO1_origin_orCat        = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO1_origin_orIdent      = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO1_ctlNum              = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO1_stVal               = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO1_q                   = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO1_t                   = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO1_ctlModel            = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO1_Oper                = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO1_Oper_ctlVal         = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO1_Oper_origin         = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO1_Oper_origin_orCat   = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO1_Oper_origin_orIdent = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO1_Oper_ctlNum         = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO1_Oper_T              = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO1_Oper_Test           = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO1_Oper_Check          = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO2                     = libiec61850go.NewDataObject()
	iedModel_GenericIO_GGIO1_SPCSO2_stVal               = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO2_q                   = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO2_Oper                = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO2_Oper_ctlVal         = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO2_Oper_origin         = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO2_Oper_origin_orCat   = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO2_Oper_origin_orIdent = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO2_Oper_ctlNum         = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO2_Oper_T              = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO2_Oper_Test           = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO2_Oper_Check          = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO2_ctlModel            = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO2_t                   = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO3                     = libiec61850go.NewDataObject()
	iedModel_GenericIO_GGIO1_SPCSO3_stVal               = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO3_q                   = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO3_Oper                = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO3_Oper_ctlVal         = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO3_Oper_origin         = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO3_Oper_origin_orCat   = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO3_Oper_origin_orIdent = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO3_Oper_ctlNum         = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO3_Oper_T              = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO3_Oper_Test           = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO3_Oper_Check          = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO3_ctlModel            = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO3_t                   = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO4                     = libiec61850go.NewDataObject()
	iedModel_GenericIO_GGIO1_SPCSO4_stVal               = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO4_q                   = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO4_Oper                = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO4_Oper_ctlVal         = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO4_Oper_origin         = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO4_Oper_origin_orCat   = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO4_Oper_origin_orIdent = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO4_Oper_ctlNum         = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO4_Oper_T              = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO4_Oper_Test           = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO4_Oper_Check          = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO4_ctlModel            = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_SPCSO4_t                   = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Ind1                       = libiec61850go.NewDataObject()
	iedModel_GenericIO_GGIO1_Ind1_stVal                 = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Ind1_q                     = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Ind1_t                     = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Ind2                       = libiec61850go.NewDataObject()
	iedModel_GenericIO_GGIO1_Ind2_stVal                 = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Ind2_q                     = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Ind2_t                     = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Ind3                       = libiec61850go.NewDataObject()
	iedModel_GenericIO_GGIO1_Ind3_stVal                 = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Ind3_q                     = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Ind3_t                     = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Ind4                       = libiec61850go.NewDataObject()
	iedModel_GenericIO_GGIO1_Ind4_stVal                 = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Ind4_q                     = libiec61850go.NewDataAttribute()
	iedModel_GenericIO_GGIO1_Ind4_t                     = libiec61850go.NewDataAttribute()

	iedModel_GenericIO_LLN0_report0 = libiec61850go.NewReportControlBlock()
	iedModel_GenericIO_LLN0_report1 = libiec61850go.NewReportControlBlock()
	iedModel_GenericIO_LLN0_report2 = libiec61850go.NewReportControlBlock()
	iedModel_GenericIO_LLN0_report3 = libiec61850go.NewReportControlBlock()
	iedModel_GenericIO_LLN0_report4 = libiec61850go.NewReportControlBlock()
	iedModel_GenericIO_LLN0_report5 = libiec61850go.NewReportControlBlock()
	iedModel_GenericIO_LLN0_report6 = libiec61850go.NewReportControlBlock()
	iedModel_GenericIO_LLN0_report7 = libiec61850go.NewReportControlBlock()
	iedModel_GenericIO_LLN0_report8 = libiec61850go.NewReportControlBlock()
	iedModel_GenericIO_LLN0_report9 = libiec61850go.NewReportControlBlock()
)

func init() {
	iedModelds_GenericIO_LLN0_Events_fcda0.Initialize(
		"GenericIO",
		false,
		"GGIO1$ST$SPCSO1$stVal",
		-1,
		nil,
		nil,
		iedModelds_GenericIO_LLN0_Events_fcda1,
	)

	iedModelds_GenericIO_LLN0_Events_fcda1.Initialize(
		"GenericIO",
		false,
		"GGIO1$ST$SPCSO2$stVal",
		-1,
		nil,
		nil,
		iedModelds_GenericIO_LLN0_Events_fcda2,
	)

	iedModelds_GenericIO_LLN0_Events_fcda2.Initialize(
		"GenericIO",
		false,
		"GGIO1$ST$SPCSO3$stVal",
		-1,
		nil,
		nil,
		iedModelds_GenericIO_LLN0_Events_fcda3,
	)

	iedModelds_GenericIO_LLN0_Events_fcda3.Initialize(
		"GenericIO",
		false,
		"GGIO1$ST$SPCSO4$stVal",
		-1,
		nil,
		nil,
		nil,
	)

	iedModelds_GenericIO_LLN0_Events.Initialize(
		"GenericIO",
		"LLN0$Events",
		4,
		iedModelds_GenericIO_LLN0_Events_fcda0,
		iedModelds_GenericIO_LLN0_Events2,
	)

	iedModelds_GenericIO_LLN0_Events2_fcda0.Initialize(
		"GenericIO",
		false,
		"GGIO1$ST$SPCSO1",
		-1,
		nil,
		nil,
		iedModelds_GenericIO_LLN0_Events2_fcda1,
	)

	iedModelds_GenericIO_LLN0_Events2_fcda1.Initialize(
		"GenericIO",
		false,
		"GGIO1$ST$SPCSO2",
		-1,
		nil,
		nil,
		iedModelds_GenericIO_LLN0_Events2_fcda2,
	)

	iedModelds_GenericIO_LLN0_Events2_fcda2.Initialize(
		"GenericIO",
		false,
		"GGIO1$ST$SPCSO3",
		-1,
		nil,
		nil,
		iedModelds_GenericIO_LLN0_Events2_fcda3,
	)

	iedModelds_GenericIO_LLN0_Events2_fcda3.Initialize(
		"GenericIO",
		false,
		"GGIO1$ST$SPCSO4",
		-1,
		nil,
		nil,
		nil,
	)

	iedModelds_GenericIO_LLN0_Events2.Initialize(
		"GenericIO",
		"LLN0$Events2",
		4,
		iedModelds_GenericIO_LLN0_Events2_fcda0,
		iedModelds_GenericIO_LLN0_Measurements,
	)

	iedModel_GenericIO.Initialize(
		libiec61850go.LogicalDeviceModelType,
		"GenericIO",
		iedModel,
		nil,
		iedModel_GenericIO_LLN0,
		nil,
	)

	iedModel_GenericIO_LLN0.Initialize(
		libiec61850go.LogicalNodeModelType,
		"LLN0",
		iedModel_GenericIO,
		iedModel_GenericIO_LPHD1,
		iedModel_GenericIO_LLN0_Mod,
	)

	iedModel_GenericIO_LLN0_Mod.Initialize(
		libiec61850go.DataObjectModelType,
		"Mod",
		iedModel_GenericIO_LLN0,
		iedModel_GenericIO_LLN0_Beh,
		iedModel_GenericIO_LLN0_Mod_stVal,
		0,
		-1,
	)

	iedModel_GenericIO_LLN0_Mod_stVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"stVal",
		iedModel_GenericIO_LLN0_Mod,
		iedModel_GenericIO_LLN0_Mod_q,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_ENUMERATED,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_LLN0_Mod_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_LLN0_Mod,
		iedModel_GenericIO_LLN0_Mod_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_LLN0_Mod_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_LLN0_Mod,
		iedModel_GenericIO_LLN0_Mod_ctlModel,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_LLN0_Mod_ctlModel.Initialize(
		libiec61850go.DataAttributeModelType,
		"ctlModel",
		iedModel_GenericIO_LLN0_Mod,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CF,
		libiec61850go.IEC61850_ENUMERATED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_LLN0_Beh.Initialize(
		libiec61850go.DataObjectModelType,
		"Beh",
		iedModel_GenericIO_LLN0,
		iedModel_GenericIO_LLN0_Health,
		iedModel_GenericIO_LLN0_Beh_stVal,
		0,
		-1,
	)

	iedModel_GenericIO_LLN0_Beh_stVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"stVal",
		iedModel_GenericIO_LLN0_Beh,
		iedModel_GenericIO_LLN0_Beh_q,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_ENUMERATED,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_LLN0_Beh_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_LLN0_Beh,
		iedModel_GenericIO_LLN0_Beh_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_LLN0_Beh_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_LLN0_Beh,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_LLN0_Health.Initialize(
		libiec61850go.DataObjectModelType,
		"Health",
		iedModel_GenericIO_LLN0,
		iedModel_GenericIO_LLN0_NamPlt,
		iedModel_GenericIO_LLN0_Health_stVal,
		0,
		-1,
	)

	iedModel_GenericIO_LLN0_Health_stVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"stVal",
		iedModel_GenericIO_LLN0_Health,
		iedModel_GenericIO_LLN0_Health_q,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_ENUMERATED,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_LLN0_Health_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_LLN0_Health,
		iedModel_GenericIO_LLN0_Health_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_LLN0_Health_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_LLN0_Health,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_LLN0_NamPlt.Initialize(
		libiec61850go.DataObjectModelType,
		"NamPlt",
		iedModel_GenericIO_LLN0,
		nil,
		iedModel_GenericIO_LLN0_NamPlt_vendor,
		0,
		-1,
	)

	iedModel_GenericIO_LLN0_NamPlt_vendor.Initialize(
		libiec61850go.DataAttributeModelType,
		"vendor",
		iedModel_GenericIO_LLN0_NamPlt,
		iedModel_GenericIO_LLN0_NamPlt_swRev,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_DC,
		libiec61850go.IEC61850_VISIBLE_STRING_255,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_LLN0_NamPlt_swRev.Initialize(
		libiec61850go.DataAttributeModelType,
		"swRev",
		iedModel_GenericIO_LLN0_NamPlt,
		iedModel_GenericIO_LLN0_NamPlt_d,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_DC,
		libiec61850go.IEC61850_VISIBLE_STRING_255,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_LLN0_NamPlt_d.Initialize(
		libiec61850go.DataAttributeModelType,
		"d",
		iedModel_GenericIO_LLN0_NamPlt,
		iedModel_GenericIO_LLN0_NamPlt_configRev,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_DC,
		libiec61850go.IEC61850_VISIBLE_STRING_255,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_LLN0_NamPlt_configRev.Initialize(
		libiec61850go.DataAttributeModelType,
		"configRev",
		iedModel_GenericIO_LLN0_NamPlt,
		iedModel_GenericIO_LLN0_NamPlt_ldNs,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_DC,
		libiec61850go.IEC61850_VISIBLE_STRING_255,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_LLN0_NamPlt_ldNs.Initialize(
		libiec61850go.DataAttributeModelType,
		"ldNs",
		iedModel_GenericIO_LLN0_NamPlt,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_EX,
		libiec61850go.IEC61850_VISIBLE_STRING_255,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_LPHD1.Initialize(
		libiec61850go.LogicalNodeModelType,
		"LPHD1",
		iedModel_GenericIO,
		iedModel_GenericIO_GGIO1,
		iedModel_GenericIO_LPHD1_PhyNam,
	)

	iedModel_GenericIO_LPHD1_PhyNam.Initialize(
		libiec61850go.DataObjectModelType,
		"PhyNam",
		iedModel_GenericIO_LPHD1,
		iedModel_GenericIO_LPHD1_PhyHealth,
		iedModel_GenericIO_LPHD1_PhyNam_vendor,
		0,
		-1,
	)

	iedModel_GenericIO_LPHD1_PhyNam_vendor.Initialize(
		libiec61850go.DataAttributeModelType,
		"vendor",
		iedModel_GenericIO_LPHD1_PhyNam,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_DC,
		libiec61850go.IEC61850_VISIBLE_STRING_255,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_LPHD1_PhyHealth.Initialize(
		libiec61850go.DataObjectModelType,
		"PhyHealth",
		iedModel_GenericIO_LPHD1,
		iedModel_GenericIO_LPHD1_Proxy,
		iedModel_GenericIO_LPHD1_PhyHealth_stVal,
		0,
		-1,
	)

	iedModel_GenericIO_LPHD1_PhyHealth_stVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"stVal",
		iedModel_GenericIO_LPHD1_PhyHealth,
		iedModel_GenericIO_LPHD1_PhyHealth_q,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_ENUMERATED,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_LPHD1_PhyHealth_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_LPHD1_PhyHealth,
		iedModel_GenericIO_LPHD1_PhyHealth_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_LPHD1_PhyHealth_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_LPHD1_PhyHealth,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_LPHD1_Proxy.Initialize(
		libiec61850go.DataObjectModelType,
		"Proxy",
		iedModel_GenericIO_LPHD1,
		nil,
		iedModel_GenericIO_LPHD1_Proxy_stVal,
		0,
		-1,
	)

	iedModel_GenericIO_LPHD1_Proxy_stVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"stVal",
		iedModel_GenericIO_LPHD1_Proxy,
		iedModel_GenericIO_LPHD1_Proxy_q,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_BOOLEAN,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_LPHD1_Proxy_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_LPHD1_Proxy,
		iedModel_GenericIO_LPHD1_Proxy_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_LPHD1_Proxy_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_LPHD1_Proxy,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1.Initialize(
		libiec61850go.LogicalNodeModelType,
		"GGIO1",
		iedModel_GenericIO,
		nil,
		iedModel_GenericIO_GGIO1_Mod,
	)

	iedModel_GenericIO_GGIO1_Mod.Initialize(
		libiec61850go.DataObjectModelType,
		"Mod",
		iedModel_GenericIO_GGIO1,
		iedModel_GenericIO_GGIO1_Beh,
		iedModel_GenericIO_GGIO1_Mod_stVal,
		0,
		-1,
	)

	iedModel_GenericIO_GGIO1_Mod_stVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"stVal",
		iedModel_GenericIO_GGIO1_Mod,
		iedModel_GenericIO_GGIO1_Mod_q,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_ENUMERATED,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Mod_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_GGIO1_Mod,
		iedModel_GenericIO_GGIO1_Mod_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Mod_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_GGIO1_Mod,
		iedModel_GenericIO_GGIO1_Mod_ctlModel,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Mod_ctlModel.Initialize(
		libiec61850go.DataAttributeModelType,
		"ctlModel",
		iedModel_GenericIO_GGIO1_Mod,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CF,
		libiec61850go.IEC61850_ENUMERATED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Beh.Initialize(
		libiec61850go.DataObjectModelType,
		"Beh",
		iedModel_GenericIO_GGIO1,
		iedModel_GenericIO_GGIO1_Health,
		iedModel_GenericIO_GGIO1_Beh_stVal,
		0,
		-1,
	)

	iedModel_GenericIO_GGIO1_Beh_stVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"stVal",
		iedModel_GenericIO_GGIO1_Beh,
		iedModel_GenericIO_GGIO1_Beh_q,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_ENUMERATED,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Beh_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_GGIO1_Beh,
		iedModel_GenericIO_GGIO1_Beh_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Beh_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_GGIO1_Beh,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Health.Initialize(
		libiec61850go.DataObjectModelType,
		"Health",
		iedModel_GenericIO_GGIO1,
		iedModel_GenericIO_GGIO1_NamPlt,
		iedModel_GenericIO_GGIO1_Health_stVal,
		0,
		-1,
	)

	iedModel_GenericIO_GGIO1_Health_stVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"stVal",
		iedModel_GenericIO_GGIO1_Health,
		iedModel_GenericIO_GGIO1_Health_q,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_ENUMERATED,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Health_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_GGIO1_Health,
		iedModel_GenericIO_GGIO1_Health_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Health_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_GGIO1_Health,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_NamPlt.Initialize(
		libiec61850go.DataObjectModelType,
		"NamPlt",
		iedModel_GenericIO_GGIO1,
		iedModel_GenericIO_GGIO1_AnIn1,
		iedModel_GenericIO_GGIO1_NamPlt_vendor,
		0,
		-1,
	)

	iedModel_GenericIO_GGIO1_NamPlt_vendor.Initialize(
		libiec61850go.DataAttributeModelType,
		"vendor",
		iedModel_GenericIO_GGIO1_NamPlt,
		iedModel_GenericIO_GGIO1_NamPlt_swRev,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_DC,
		libiec61850go.IEC61850_VISIBLE_STRING_255,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_NamPlt_swRev.Initialize(
		libiec61850go.DataAttributeModelType,
		"swRev",
		iedModel_GenericIO_GGIO1_NamPlt,
		iedModel_GenericIO_GGIO1_NamPlt_d,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_DC,
		libiec61850go.IEC61850_VISIBLE_STRING_255,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_NamPlt_d.Initialize(
		libiec61850go.DataAttributeModelType,
		"d",
		iedModel_GenericIO_GGIO1_NamPlt,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_DC,
		libiec61850go.IEC61850_VISIBLE_STRING_255,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_AnIn1.Initialize(
		libiec61850go.DataObjectModelType,
		"AnIn1",
		iedModel_GenericIO_GGIO1,
		iedModel_GenericIO_GGIO1_AnIn2,
		iedModel_GenericIO_GGIO1_AnIn1_mag,
		0,
		-1,
	)

	iedModel_GenericIO_GGIO1_AnIn1_mag.Initialize(
		libiec61850go.DataAttributeModelType,
		"mag",
		iedModel_GenericIO_GGIO1_AnIn1,
		iedModel_GenericIO_GGIO1_AnIn1_q,
		iedModel_GenericIO_GGIO1_AnIn1_mag_f,
		0,
		-1,
		libiec61850go.IEC61850_FC_MX,
		libiec61850go.IEC61850_CONSTRUCTED,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_AnIn1_mag_f.Initialize(
		libiec61850go.DataAttributeModelType,
		"f",
		iedModel_GenericIO_GGIO1_AnIn1_mag,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_MX,
		libiec61850go.IEC61850_FLOAT32,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_AnIn1_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_GGIO1_AnIn1,
		iedModel_GenericIO_GGIO1_AnIn1_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_MX,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_AnIn1_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_GGIO1_AnIn1,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_MX,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_AnIn2.Initialize(
		libiec61850go.DataObjectModelType,
		"AnIn2",
		iedModel_GenericIO_GGIO1,
		iedModel_GenericIO_GGIO1_AnIn3,
		iedModel_GenericIO_GGIO1_AnIn2_mag,
		0,
		-1,
	)

	iedModel_GenericIO_GGIO1_AnIn2_mag.Initialize(
		libiec61850go.DataAttributeModelType,
		"mag",
		iedModel_GenericIO_GGIO1_AnIn2,
		iedModel_GenericIO_GGIO1_AnIn2_q,
		iedModel_GenericIO_GGIO1_AnIn2_mag_f,
		0,
		-1,
		libiec61850go.IEC61850_FC_MX,
		libiec61850go.IEC61850_CONSTRUCTED,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_AnIn2_mag_f.Initialize(
		libiec61850go.DataAttributeModelType,
		"f",
		iedModel_GenericIO_GGIO1_AnIn2_mag,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_MX,
		libiec61850go.IEC61850_FLOAT32,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_AnIn2_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_GGIO1_AnIn2,
		iedModel_GenericIO_GGIO1_AnIn2_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_MX,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_AnIn2_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_GGIO1_AnIn2,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_MX,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_AnIn3.Initialize(
		libiec61850go.DataObjectModelType,
		"AnIn3",
		iedModel_GenericIO_GGIO1,
		iedModel_GenericIO_GGIO1_AnIn4,
		iedModel_GenericIO_GGIO1_AnIn3_mag,
		0,
		-1,
	)

	iedModel_GenericIO_GGIO1_AnIn3_mag.Initialize(
		libiec61850go.DataAttributeModelType,
		"mag",
		iedModel_GenericIO_GGIO1_AnIn3,
		iedModel_GenericIO_GGIO1_AnIn3_q,
		iedModel_GenericIO_GGIO1_AnIn3_mag_f,
		0,
		-1,
		libiec61850go.IEC61850_FC_MX,
		libiec61850go.IEC61850_CONSTRUCTED,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_AnIn3_mag_f.Initialize(
		libiec61850go.DataAttributeModelType,
		"f",
		iedModel_GenericIO_GGIO1_AnIn3_mag,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_MX,
		libiec61850go.IEC61850_FLOAT32,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_AnIn3_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_GGIO1_AnIn3,
		iedModel_GenericIO_GGIO1_AnIn3_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_MX,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_AnIn3_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_GGIO1_AnIn3,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_MX,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_AnIn4.Initialize(
		libiec61850go.DataObjectModelType,
		"AnIn4",
		iedModel_GenericIO_GGIO1,
		iedModel_GenericIO_GGIO1_SPCSO1,
		iedModel_GenericIO_GGIO1_AnIn4_mag,
		0,
		-1,
	)

	iedModel_GenericIO_GGIO1_AnIn4_mag.Initialize(
		libiec61850go.DataAttributeModelType,
		"mag",
		iedModel_GenericIO_GGIO1_AnIn4,
		iedModel_GenericIO_GGIO1_AnIn4_q,
		iedModel_GenericIO_GGIO1_AnIn4_mag_f,
		0,
		-1,
		libiec61850go.IEC61850_FC_MX,
		libiec61850go.IEC61850_CONSTRUCTED,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_AnIn4_mag_f.Initialize(
		libiec61850go.DataAttributeModelType,
		"f",
		iedModel_GenericIO_GGIO1_AnIn4_mag,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_MX,
		libiec61850go.IEC61850_FLOAT32,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_AnIn4_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_GGIO1_AnIn4,
		iedModel_GenericIO_GGIO1_AnIn4_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_MX,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_AnIn4_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_GGIO1_AnIn4,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_MX,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO1.Initialize(
		libiec61850go.DataObjectModelType,
		"SPCSO1",
		iedModel_GenericIO_GGIO1,
		iedModel_GenericIO_GGIO1_SPCSO2,
		iedModel_GenericIO_GGIO1_SPCSO1_origin,
		0,
		-1,
	)

	iedModel_GenericIO_GGIO1_SPCSO1_origin.Initialize(
		libiec61850go.DataAttributeModelType,
		"origin",
		iedModel_GenericIO_GGIO1_SPCSO1,
		iedModel_GenericIO_GGIO1_SPCSO1_ctlNum,
		iedModel_GenericIO_GGIO1_SPCSO1_origin_orCat,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_CONSTRUCTED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO1_origin_orCat.Initialize(
		libiec61850go.DataAttributeModelType,
		"orCat",
		iedModel_GenericIO_GGIO1_SPCSO1_origin,
		iedModel_GenericIO_GGIO1_SPCSO1_origin_orIdent,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_ENUMERATED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO1_origin_orIdent.Initialize(
		libiec61850go.DataAttributeModelType,
		"orIdent",
		iedModel_GenericIO_GGIO1_SPCSO1_origin,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_OCTET_STRING_64,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO1_ctlNum.Initialize(
		libiec61850go.DataAttributeModelType,
		"ctlNum",
		iedModel_GenericIO_GGIO1_SPCSO1,
		iedModel_GenericIO_GGIO1_SPCSO1_stVal,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_INT8U,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO1_stVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"stVal",
		iedModel_GenericIO_GGIO1_SPCSO1,
		iedModel_GenericIO_GGIO1_SPCSO1_q,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_BOOLEAN,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO1_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_GGIO1_SPCSO1,
		iedModel_GenericIO_GGIO1_SPCSO1_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO1_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_GGIO1_SPCSO1,
		iedModel_GenericIO_GGIO1_SPCSO1_ctlModel,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO1_ctlModel.Initialize(
		libiec61850go.DataAttributeModelType,
		"ctlModel",
		iedModel_GenericIO_GGIO1_SPCSO1,
		iedModel_GenericIO_GGIO1_SPCSO1_Oper,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CF,
		libiec61850go.IEC61850_ENUMERATED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO1_Oper.Initialize(
		libiec61850go.DataAttributeModelType,
		"Oper",
		iedModel_GenericIO_GGIO1_SPCSO1,
		nil,
		iedModel_GenericIO_GGIO1_SPCSO1_Oper_ctlVal,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_CONSTRUCTED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO1_Oper_ctlVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"ctlVal",
		iedModel_GenericIO_GGIO1_SPCSO1_Oper,
		iedModel_GenericIO_GGIO1_SPCSO1_Oper_origin,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_BOOLEAN,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO1_Oper_origin.Initialize(
		libiec61850go.DataAttributeModelType,
		"origin",
		iedModel_GenericIO_GGIO1_SPCSO1_Oper,
		iedModel_GenericIO_GGIO1_SPCSO1_Oper_ctlNum,
		iedModel_GenericIO_GGIO1_SPCSO1_Oper_origin_orCat,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_CONSTRUCTED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO1_Oper_origin_orCat.Initialize(
		libiec61850go.DataAttributeModelType,
		"orCat",
		iedModel_GenericIO_GGIO1_SPCSO1_Oper_origin,
		iedModel_GenericIO_GGIO1_SPCSO1_Oper_origin_orIdent,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_ENUMERATED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO1_Oper_origin_orIdent.Initialize(
		libiec61850go.DataAttributeModelType,
		"orIdent",
		iedModel_GenericIO_GGIO1_SPCSO1_Oper_origin,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_OCTET_STRING_64,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO1_Oper_ctlNum.Initialize(
		libiec61850go.DataAttributeModelType,
		"ctlNum",
		iedModel_GenericIO_GGIO1_SPCSO1_Oper,
		iedModel_GenericIO_GGIO1_SPCSO1_Oper_T,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_INT8U,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO1_Oper_T.Initialize(
		libiec61850go.DataAttributeModelType,
		"T",
		iedModel_GenericIO_GGIO1_SPCSO1_Oper,
		iedModel_GenericIO_GGIO1_SPCSO1_Oper_Test,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO1_Oper_Test.Initialize(
		libiec61850go.DataAttributeModelType,
		"Test",
		iedModel_GenericIO_GGIO1_SPCSO1_Oper,
		iedModel_GenericIO_GGIO1_SPCSO1_Oper_Check,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_BOOLEAN,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO1_Oper_Check.Initialize(
		libiec61850go.DataAttributeModelType,
		"Check",
		iedModel_GenericIO_GGIO1_SPCSO1_Oper,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_CHECK,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO2.Initialize(
		libiec61850go.DataObjectModelType,
		"SPCSO2",
		iedModel_GenericIO_GGIO1,
		iedModel_GenericIO_GGIO1_SPCSO3,
		iedModel_GenericIO_GGIO1_SPCSO2_stVal,
		0,
		-1,
	)

	iedModel_GenericIO_GGIO1_SPCSO2_stVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"stVal",
		iedModel_GenericIO_GGIO1_SPCSO2,
		iedModel_GenericIO_GGIO1_SPCSO2_q,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_BOOLEAN,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO2_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_GGIO1_SPCSO2,
		iedModel_GenericIO_GGIO1_SPCSO2_Oper,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO2_Oper.Initialize(
		libiec61850go.DataAttributeModelType,
		"Oper",
		iedModel_GenericIO_GGIO1_SPCSO2,
		iedModel_GenericIO_GGIO1_SPCSO2_ctlModel,
		iedModel_GenericIO_GGIO1_SPCSO2_Oper_ctlVal,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_CONSTRUCTED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO2_Oper_ctlVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"ctlVal",
		iedModel_GenericIO_GGIO1_SPCSO2_Oper,
		iedModel_GenericIO_GGIO1_SPCSO2_Oper_origin,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_BOOLEAN,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO2_Oper_origin.Initialize(
		libiec61850go.DataAttributeModelType,
		"origin",
		iedModel_GenericIO_GGIO1_SPCSO2_Oper,
		iedModel_GenericIO_GGIO1_SPCSO2_Oper_ctlNum,
		iedModel_GenericIO_GGIO1_SPCSO2_Oper_origin_orCat,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_CONSTRUCTED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO2_Oper_origin_orCat.Initialize(
		libiec61850go.DataAttributeModelType,
		"orCat",
		iedModel_GenericIO_GGIO1_SPCSO2_Oper_origin,
		iedModel_GenericIO_GGIO1_SPCSO2_Oper_origin_orIdent,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_ENUMERATED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO2_Oper_origin_orIdent.Initialize(
		libiec61850go.DataAttributeModelType,
		"orIdent",
		iedModel_GenericIO_GGIO1_SPCSO2_Oper_origin,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_OCTET_STRING_64,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO2_Oper_ctlNum.Initialize(
		libiec61850go.DataAttributeModelType,
		"ctlNum",
		iedModel_GenericIO_GGIO1_SPCSO2_Oper,
		iedModel_GenericIO_GGIO1_SPCSO2_Oper_T,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_INT8U,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO2_Oper_T.Initialize(
		libiec61850go.DataAttributeModelType,
		"T",
		iedModel_GenericIO_GGIO1_SPCSO2_Oper,
		iedModel_GenericIO_GGIO1_SPCSO2_Oper_Test,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO2_Oper_Test.Initialize(
		libiec61850go.DataAttributeModelType,
		"Test",
		iedModel_GenericIO_GGIO1_SPCSO2_Oper,
		iedModel_GenericIO_GGIO1_SPCSO2_Oper_Check,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_BOOLEAN,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO2_Oper_Check.Initialize(
		libiec61850go.DataAttributeModelType,
		"Check",
		iedModel_GenericIO_GGIO1_SPCSO2_Oper,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_CHECK,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO2_ctlModel.Initialize(
		libiec61850go.DataAttributeModelType,
		"ctlModel",
		iedModel_GenericIO_GGIO1_SPCSO2,
		iedModel_GenericIO_GGIO1_SPCSO2_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CF,
		libiec61850go.IEC61850_ENUMERATED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO2_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_GGIO1_SPCSO2,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO3.Initialize(
		libiec61850go.DataObjectModelType,
		"SPCSO3",
		iedModel_GenericIO_GGIO1,
		iedModel_GenericIO_GGIO1_SPCSO4,
		iedModel_GenericIO_GGIO1_SPCSO3_stVal,
		0,
		-1,
	)

	iedModel_GenericIO_GGIO1_SPCSO3_stVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"stVal",
		iedModel_GenericIO_GGIO1_SPCSO3,
		iedModel_GenericIO_GGIO1_SPCSO3_q,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_BOOLEAN,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO3_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_GGIO1_SPCSO3,
		iedModel_GenericIO_GGIO1_SPCSO3_Oper,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO3_Oper.Initialize(
		libiec61850go.DataAttributeModelType,
		"Oper",
		iedModel_GenericIO_GGIO1_SPCSO3,
		iedModel_GenericIO_GGIO1_SPCSO3_ctlModel,
		iedModel_GenericIO_GGIO1_SPCSO3_Oper_ctlVal,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_CONSTRUCTED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO3_Oper_ctlVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"ctlVal",
		iedModel_GenericIO_GGIO1_SPCSO3_Oper,
		iedModel_GenericIO_GGIO1_SPCSO3_Oper_origin,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_BOOLEAN,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO3_Oper_origin.Initialize(
		libiec61850go.DataAttributeModelType,
		"origin",
		iedModel_GenericIO_GGIO1_SPCSO3_Oper,
		iedModel_GenericIO_GGIO1_SPCSO3_Oper_ctlNum,
		iedModel_GenericIO_GGIO1_SPCSO3_Oper_origin_orCat,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_CONSTRUCTED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO3_Oper_origin_orCat.Initialize(
		libiec61850go.DataAttributeModelType,
		"orCat",
		iedModel_GenericIO_GGIO1_SPCSO3_Oper_origin,
		iedModel_GenericIO_GGIO1_SPCSO3_Oper_origin_orIdent,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_ENUMERATED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO3_Oper_origin_orIdent.Initialize(
		libiec61850go.DataAttributeModelType,
		"orIdent",
		iedModel_GenericIO_GGIO1_SPCSO3_Oper_origin,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_OCTET_STRING_64,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO3_Oper_ctlNum.Initialize(
		libiec61850go.DataAttributeModelType,
		"ctlNum",
		iedModel_GenericIO_GGIO1_SPCSO3_Oper,
		iedModel_GenericIO_GGIO1_SPCSO3_Oper_T,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_INT8U,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO3_Oper_T.Initialize(
		libiec61850go.DataAttributeModelType,
		"T",
		iedModel_GenericIO_GGIO1_SPCSO3_Oper,
		iedModel_GenericIO_GGIO1_SPCSO3_Oper_Test,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO3_Oper_Test.Initialize(
		libiec61850go.DataAttributeModelType,
		"Test",
		iedModel_GenericIO_GGIO1_SPCSO3_Oper,
		iedModel_GenericIO_GGIO1_SPCSO3_Oper_Check,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_BOOLEAN,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO3_Oper_Check.Initialize(
		libiec61850go.DataAttributeModelType,
		"Check",
		iedModel_GenericIO_GGIO1_SPCSO3_Oper,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_CHECK,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO3_ctlModel.Initialize(
		libiec61850go.DataAttributeModelType,
		"ctlModel",
		iedModel_GenericIO_GGIO1_SPCSO3,
		iedModel_GenericIO_GGIO1_SPCSO3_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CF,
		libiec61850go.IEC61850_ENUMERATED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO3_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_GGIO1_SPCSO3,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO4.Initialize(
		libiec61850go.DataObjectModelType,
		"SPCSO4",
		iedModel_GenericIO_GGIO1,
		iedModel_GenericIO_GGIO1_Ind1,
		iedModel_GenericIO_GGIO1_SPCSO4_stVal,
		0,
		-1,
	)

	iedModel_GenericIO_GGIO1_SPCSO4_stVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"stVal",
		iedModel_GenericIO_GGIO1_SPCSO4,
		iedModel_GenericIO_GGIO1_SPCSO4_q,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_BOOLEAN,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO4_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_GGIO1_SPCSO4,
		iedModel_GenericIO_GGIO1_SPCSO4_Oper,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO4_Oper.Initialize(
		libiec61850go.DataAttributeModelType,
		"Oper",
		iedModel_GenericIO_GGIO1_SPCSO4,
		iedModel_GenericIO_GGIO1_SPCSO4_ctlModel,
		iedModel_GenericIO_GGIO1_SPCSO4_Oper_ctlVal,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_CONSTRUCTED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO4_Oper_ctlVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"ctlVal",
		iedModel_GenericIO_GGIO1_SPCSO4_Oper,
		iedModel_GenericIO_GGIO1_SPCSO4_Oper_origin,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_BOOLEAN,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO4_Oper_origin.Initialize(
		libiec61850go.DataAttributeModelType,
		"origin",
		iedModel_GenericIO_GGIO1_SPCSO4_Oper,
		iedModel_GenericIO_GGIO1_SPCSO4_Oper_ctlNum,
		iedModel_GenericIO_GGIO1_SPCSO4_Oper_origin_orCat,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_CONSTRUCTED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO4_Oper_origin_orCat.Initialize(
		libiec61850go.DataAttributeModelType,
		"orCat",
		iedModel_GenericIO_GGIO1_SPCSO4_Oper_origin,
		iedModel_GenericIO_GGIO1_SPCSO4_Oper_origin_orIdent,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_ENUMERATED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO4_Oper_origin_orIdent.Initialize(
		libiec61850go.DataAttributeModelType,
		"orIdent",
		iedModel_GenericIO_GGIO1_SPCSO4_Oper_origin,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_OCTET_STRING_64,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO4_Oper_ctlNum.Initialize(
		libiec61850go.DataAttributeModelType,
		"ctlNum",
		iedModel_GenericIO_GGIO1_SPCSO4_Oper,
		iedModel_GenericIO_GGIO1_SPCSO4_Oper_T,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_INT8U,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO4_Oper_T.Initialize(
		libiec61850go.DataAttributeModelType,
		"T",
		iedModel_GenericIO_GGIO1_SPCSO4_Oper,
		iedModel_GenericIO_GGIO1_SPCSO4_Oper_Test,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO4_Oper_Test.Initialize(
		libiec61850go.DataAttributeModelType,
		"Test",
		iedModel_GenericIO_GGIO1_SPCSO4_Oper,
		iedModel_GenericIO_GGIO1_SPCSO4_Oper_Check,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_BOOLEAN,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO4_Oper_Check.Initialize(
		libiec61850go.DataAttributeModelType,
		"Check",
		iedModel_GenericIO_GGIO1_SPCSO4_Oper,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CO,
		libiec61850go.IEC61850_CHECK,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO4_ctlModel.Initialize(
		libiec61850go.DataAttributeModelType,
		"ctlModel",
		iedModel_GenericIO_GGIO1_SPCSO4,
		iedModel_GenericIO_GGIO1_SPCSO4_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_CF,
		libiec61850go.IEC61850_ENUMERATED,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_SPCSO4_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_GGIO1_SPCSO4,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Ind1.Initialize(
		libiec61850go.DataObjectModelType,
		"Ind1",
		iedModel_GenericIO_GGIO1,
		iedModel_GenericIO_GGIO1_Ind2,
		iedModel_GenericIO_GGIO1_Ind1_stVal,
		0,
		-1,
	)

	iedModel_GenericIO_GGIO1_Ind1_stVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"stVal",
		iedModel_GenericIO_GGIO1_Ind1,
		iedModel_GenericIO_GGIO1_Ind1_q,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_BOOLEAN,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Ind1_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_GGIO1_Ind1,
		iedModel_GenericIO_GGIO1_Ind1_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Ind1_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_GGIO1_Ind1,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Ind2.Initialize(
		libiec61850go.DataObjectModelType,
		"Ind2",
		iedModel_GenericIO_GGIO1,
		iedModel_GenericIO_GGIO1_Ind3,
		iedModel_GenericIO_GGIO1_Ind2_stVal,
		0,
		-1,
	)

	iedModel_GenericIO_GGIO1_Ind2_stVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"stVal",
		iedModel_GenericIO_GGIO1_Ind2,
		iedModel_GenericIO_GGIO1_Ind2_q,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_BOOLEAN,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Ind2_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_GGIO1_Ind2,
		iedModel_GenericIO_GGIO1_Ind2_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Ind2_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_GGIO1_Ind2,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Ind3.Initialize(
		libiec61850go.DataObjectModelType,
		"Ind3",
		iedModel_GenericIO_GGIO1,
		iedModel_GenericIO_GGIO1_Ind4,
		iedModel_GenericIO_GGIO1_Ind3_stVal,
		0,
		-1,
	)

	iedModel_GenericIO_GGIO1_Ind3_stVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"stVal",
		iedModel_GenericIO_GGIO1_Ind3,
		iedModel_GenericIO_GGIO1_Ind3_q,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_BOOLEAN,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Ind3_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_GGIO1_Ind3,
		iedModel_GenericIO_GGIO1_Ind3_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Ind3_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_GGIO1_Ind3,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Ind4.Initialize(
		libiec61850go.DataObjectModelType,
		"Ind4",
		iedModel_GenericIO_GGIO1,
		nil,
		iedModel_GenericIO_GGIO1_Ind4_stVal,
		0,
		-1,
	)

	iedModel_GenericIO_GGIO1_Ind4_stVal.Initialize(
		libiec61850go.DataAttributeModelType,
		"stVal",
		iedModel_GenericIO_GGIO1_Ind4,
		iedModel_GenericIO_GGIO1_Ind4_q,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_BOOLEAN,
		0+libiec61850go.TRG_OPT_DATA_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Ind4_q.Initialize(
		libiec61850go.DataAttributeModelType,
		"q",
		iedModel_GenericIO_GGIO1_Ind4,
		iedModel_GenericIO_GGIO1_Ind4_t,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_QUALITY,
		0+libiec61850go.TRG_OPT_QUALITY_CHANGED,
		nil,
		0,
	)

	iedModel_GenericIO_GGIO1_Ind4_t.Initialize(
		libiec61850go.DataAttributeModelType,
		"t",
		iedModel_GenericIO_GGIO1_Ind4,
		nil,
		nil,
		0,
		-1,
		libiec61850go.IEC61850_FC_ST,
		libiec61850go.IEC61850_TIMESTAMP,
		0,
		nil,
		0,
	)

	iedModel.Initialize(
		"simpleIO",
		iedModel_GenericIO,
		iedModelds_GenericIO_LLN0_Events,
		iedModel_GenericIO_LLN0_report0,
		nil,
		nil,
		nil,
		nil,
		nil,
		initializeValues,
	)

	iedModel_GenericIO_LLN0_report0.Initialize(iedModel_GenericIO_LLN0, "EventsRCB01", "Events1", false, "Events", 1, 88, 175, 50, 1000, []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, iedModel_GenericIO_LLN0_report1)
	iedModel_GenericIO_LLN0_report1.Initialize(iedModel_GenericIO_LLN0, "EventsRCBPreConf01", "Events1", false, "Events", 1, 88, 175, 50, 1000, []byte{0x4, 0xc0, 0xa8, 0x2, 0x9, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, iedModel_GenericIO_LLN0_report2)
	iedModel_GenericIO_LLN0_report2.Initialize(iedModel_GenericIO_LLN0, "EventsBRCB01", "Events2", true, "Events", 1, 88, 175, 50, 1000, []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, iedModel_GenericIO_LLN0_report3)
	iedModel_GenericIO_LLN0_report3.Initialize(iedModel_GenericIO_LLN0, "EventsBRCBPreConf01", "Events2", true, "Events", 1, 88, 175, 50, 1000, []byte{0x4, 0xc0, 0xa8, 0x2, 0x9, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, iedModel_GenericIO_LLN0_report4)
	iedModel_GenericIO_LLN0_report4.Initialize(iedModel_GenericIO_LLN0, "EventsIndexed01", "Events2", false, "Events", 1, 88, 175, 50, 1000, []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, iedModel_GenericIO_LLN0_report5)
	iedModel_GenericIO_LLN0_report5.Initialize(iedModel_GenericIO_LLN0, "EventsIndexed02", "Events2", false, "Events", 1, 88, 175, 50, 1000, []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, iedModel_GenericIO_LLN0_report6)
	iedModel_GenericIO_LLN0_report6.Initialize(iedModel_GenericIO_LLN0, "EventsIndexed03", "Events2", false, "Events", 1, 88, 175, 50, 1000, []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, iedModel_GenericIO_LLN0_report7)
	iedModel_GenericIO_LLN0_report7.Initialize(iedModel_GenericIO_LLN0, "Measurements01", "Measurements", true, "Measurements", 1, 80, 239, 50, 1000, []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, iedModel_GenericIO_LLN0_report8)
	iedModel_GenericIO_LLN0_report8.Initialize(iedModel_GenericIO_LLN0, "Measurements02", "Measurements", true, "Measurements", 1, 80, 239, 50, 1000, []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, iedModel_GenericIO_LLN0_report9)
	iedModel_GenericIO_LLN0_report9.Initialize(iedModel_GenericIO_LLN0, "Measurements03", "Measurements", true, "Measurements", 1, 80, 239, 50, 1000, []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, nil)
}

func initializeValues() {
	// Set initial values for the model
	fmt.Println("Initializing values for the model")
}
